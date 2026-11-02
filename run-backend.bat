@echo off
echo ========================================================
echo Menjalankan Backend API Pondok Pesantren Al-Hamid (Go)
echo ========================================================
cd backend
if exist server.exe (
    server.exe
) else (
    go run ./cmd/server
)
pause

