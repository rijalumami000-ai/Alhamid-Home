#!/usr/bin/env bash
# ==============================================================================
# Script Otomatis Deploy Production untuk: alhamidcintamulya.com
# Sistem: Ubuntu / Debian
# ==============================================================================

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}================================================================${NC}"
echo -e "${GREEN}  Memulai Proses Setup Production: alhamidcintamulya.com${NC}"
echo -e "${BLUE}================================================================${NC}"

APP_DIR="/var/www/alhamidcintamulya.com"

# 1. Pastikan Direktori Web & Log
echo -e "\n${YELLOW}[1/6] Menyiapkan direktori aplikasi dan izin akses...${NC}"
sudo mkdir -p $APP_DIR/backend/uploads
sudo mkdir -p $APP_DIR/logs
sudo chown -R $USER:$USER $APP_DIR
sudo chmod -R 755 $APP_DIR/backend/uploads

# 2. Setup Database PostgreSQL
echo -e "\n${YELLOW}[2/6] Mengonfigurasi PostgreSQL Production...${NC}"
DB_NAME="alhamid_db"
DB_USER="alhamid_user"
DB_PASS="AlhamidCintaMulya2026!Secure"

sudo -u postgres psql -tc "SELECT 1 FROM pg_database WHERE datname = '$DB_NAME'" | grep -q 1 || \
sudo -u postgres psql -c "CREATE DATABASE $DB_NAME;"

sudo -u postgres psql -tc "SELECT 1 FROM pg_roles WHERE rolname = '$DB_USER'" | grep -q 1 || \
sudo -u postgres psql -c "CREATE USER $DB_USER WITH ENCRYPTED PASSWORD '$DB_PASS';"

sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;"
sudo -u postgres psql -c "ALTER DATABASE $DB_NAME OWNER TO $DB_USER;"

echo -e "${GREEN}✓ Database $DB_NAME dan user $DB_USER siap!${NC}"

# 3. Environment Backend & Frontend
echo -e "\n${YELLOW}[3/6] Menyiapkan environment configuration...${NC}"
if [ ! -f "$APP_DIR/backend/.env" ]; then
    cat <<EOF > "$APP_DIR/backend/.env"
PORT=8088
ENV=production
DB_HOST=127.0.0.1
DB_PORT=5432
DB_USER=alhamid_user
DB_PASSWORD=AlhamidCintaMulya2026!Secure
DB_NAME=alhamid_db
DB_SSLMODE=disable
JWT_SECRET=alhamid_jwt_secret_key_2026_super_secure_random
FRONTEND_URL=https://alhamidcintamulya.com
EOF
    echo -e "${GREEN}✓ backend/.env berhasil dibuat.${NC}"
fi

if [ ! -f "$APP_DIR/frontend/.env" ]; then
    cat <<EOF > "$APP_DIR/frontend/.env"
ORIGIN=https://alhamidcintamulya.com
PORT=3005
HOST=127.0.0.1
EOF
    echo -e "${GREEN}✓ frontend/.env berhasil dibuat.${NC}"
fi

# 4. Build Frontend & Siapkan Backend Binary
echo -e "\n${YELLOW}[4/6] Mengompilasi aplikasi Frontend dan Backend...${NC}"
cd $APP_DIR/frontend
npm ci
npm run build

cd $APP_DIR/backend
if [ -f "server_linux" ]; then
    chmod +x server_linux
    cp -f server_linux server
fi
chmod +x $APP_DIR/backend/server

# 5. Jalankan Layanan dengan PM2
echo -e "\n${YELLOW}[5/6] Menjalankan layanan dengan PM2...${NC}"
cd $APP_DIR
pm2 start ecosystem.config.cjs
pm2 save

# 6. Konfigurasi Nginx
echo -e "\n${YELLOW}[6/6] Menyiapkan Nginx Reverse Proxy...${NC}"
sudo cp $APP_DIR/nginx.conf /etc/nginx/sites-available/alhamidcintamulya.com
sudo ln -sf /etc/nginx/sites-available/alhamidcintamulya.com /etc/nginx/sites-enabled/

sudo nginx -t
sudo systemctl reload nginx

echo -e "\n${GREEN}================================================================${NC}"
echo -e "${GREEN}  DEPLOYMENT SELESAI DENGAN SUKSES!${NC}"
echo -e "${GREEN}================================================================${NC}"
