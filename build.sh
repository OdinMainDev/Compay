#!/bin/bash
#
# Сборка образа в Docker для проекта Compay
#
# Использование:
#   ./build.sh <path/to/main.go>  # Собрать произвольное приложение
#

set -e

# Определение приложения для сборки (только полный путь)
APP_INPUT="${1:-}"

# Функция сборки одного приложения
build_app() {
    local app_name="$1"
    local app_path="$2"
    # Получить актуальное имя тега для назначения имени версии
    version=$(git rev-list --count HEAD 2>/dev/null || true)

    echo ""
    echo "=========================================="
    echo "Сборка: $app_name"
    echo "Путь: $app_path"
    echo "Version: $version"
    echo "=========================================="

    # Если имени тега нет, имя версии будет сформировано из предыдущего тега и суффикса -dev
    if [[ -z $version ]]; then
        lasttag=$(git tag --sort=-version:refname | head -n 1)
        if [[ $lasttag ]]; then
            lasttag="${lasttag}-"
        fi
        version="${lasttag}dev"
    fi

    echo "Версия образа: $version"

    local image_name="compay"

    docker buildx build --platform linux/amd64 --load \
        --build-arg VERSION="$version" \
        --build-arg APP_PATH="$app_path" \
        -t "${image_name}:latest" \
        -t "${image_name}:${version}" \
        .

    mkdir -p ./build
    local tar_file="./build/${image_name}-${version}.tar"
    rm -f "$tar_file"
    docker save -o "$tar_file" "${image_name}:${version}"
    chmod 755 "$tar_file"

    echo "Образ создан: ${image_name}:${version}"
    echo "Файл: $tar_file"
}


if [[ -z "$APP_INPUT" ]]; then
    echo "Нужно указать путь к main.go, например: ./build.sh ./cmd/merchant/main.go" >&2
    exit 1
fi

if [[ -f "$APP_INPUT" ]]; then
    app_name=$(basename "$(dirname "$APP_INPUT")")
    build_app "$app_name" "$APP_INPUT"
else
    echo "Путь не найден: $APP_INPUT" >&2
    exit 1
fi
