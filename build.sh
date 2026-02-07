#!/bin/bash
#
# Сборка образа в Docker для проекта Compay
#
# Использование:
#   ./build.sh                    # Собрать merchant (по умолчанию)
#   ./build.sh merchant           # Собрать merchant
#   ./build.sh btc                # Собрать btc
#   ./build.sh rates              # Собрать rates
#   ./build.sh all                # Собрать все приложения
#   ./build.sh <path/to/main.go>  # Собрать произвольное приложение
#

set -e

# Определение приложения для сборки
APP_INPUT="${1:-merchant}"

# Функция сборки одного приложения
build_app() {
    local app_name="$1"
    local app_path="$2"

    echo ""
    echo "=========================================="
    echo "Сборка: $app_name"
    echo "Путь: $app_path"
    echo "=========================================="

    # Получить актуальное имя тега для назначения имени версии
    version=$(git tag --points-at HEAD)

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

    docker build --platform linux/amd64 \
        --build-arg VERSION="$version" \
        --build-arg APP_PATH="$app_path" \
        --build-arg APP_NAME="$app_name" \
        -t "${image_name}:${version}" \
        -t "${image_name}:latest" \
        .

    mkdir -p ./build
    local tar_file="./build/${image_name}-${version}.tar"
    rm -f "$tar_file"
    docker save -o "$tar_file" "${image_name}:${version}"
    chmod 755 "$tar_file"

    echo "Образ создан: ${image_name}:${version}"
    echo "Файл: $tar_file"
}


app_name=$(basename "$(dirname "$APP_INPUT")")
build_app "$app_name" "$APP_INPUT"
