#!/bin/bash

echo "=== Verificando instalación de Xdebug ==="
echo ""

echo "1. Versión de PHP:"
docker compose exec php php -v
echo ""

echo "2. Módulos PHP instalados (buscando xdebug):"
docker compose exec php php -m | grep -i xdebug
echo ""

echo "3. Archivos de configuración en conf.d/:"
docker compose exec php ls -la /usr/local/etc/php/conf.d/
echo ""

echo "4. Contenido de docker-php-ext-xdebug.ini:"
docker compose exec php cat /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini 2>/dev/null || echo "Archivo no encontrado"
echo ""

echo "5. Extensiones compiladas:"
docker compose exec php ls -la /usr/local/lib/php/extensions/
echo ""

echo "6. Info de Xdebug (si está instalado):"
docker compose exec php php --ri xdebug 2>/dev/null || echo "Xdebug no está instalado"
