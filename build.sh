#!/bin/bash

echo "Сборка образа в Docker для проекта Compay"

# Получить актуальное имя тега для назначения имени версии
version=$(git tag --points-at HEAD)

# Если имени тега нет, имя версии будет сформировано из предыдущего тега и суффикса -dev
if [[ -z $version ]]; then
  lasttag=$(git tag --sort=-version:refname | head -n 1)
  if [[ $lasttag ]]; then
    lasttag=lasttag+"-"
  fi
  version=${lasttag}dev
fi

echo "Версия образа: $version"

docker build --platform linux/amd64 --build-arg VERSION=$version -t compay:$version .
mkdir -p ./build
rm ./build/compay-$version.tar
docker save -o ./build/compay-$version.tar compay:$version
chmod 755 ./build/compay-$version.tar

echo "Образ создан в build/compay-$version.tar"
