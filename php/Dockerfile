FROM php:8.2-cli

# Install system dependencies and build tools
RUN apt-get update && apt-get install -y \
    git \
    unzip \
    libzip-dev \
    autoconf \
    g++ \
    make \
    && docker-php-ext-install zip \
    && rm -rf /var/lib/apt/lists/*

# Install Xdebug
RUN pecl install xdebug-3.3.1 && \
    docker-php-ext-enable xdebug && \
    php -v

# Install Composer
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

# Set working directory
WORKDIR /app

# Copy Xdebug configuration (after WORKDIR so path is correct)
COPY docker/xdebug.ini /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini

# Copy composer files
COPY composer.json composer.lock* ./

# Install PHP dependencies
RUN composer install --no-interaction --no-scripts --no-autoloader --prefer-dist

# Copy application files
COPY . .

# Generate autoloader
RUN composer dump-autoload --optimize

# Default command
CMD ["php", "-a"]
