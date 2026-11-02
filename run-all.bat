@echo off
echo ========================================================
echo Membuka Backend (Port 8088) dan Frontend (Port 5173)...
echo ========================================================
start "Al-Hamid Backend (Go :8088)" cmd /k "cd backend && server.exe"
start "Al-Hamid Frontend (Svelte :5173)" cmd /k "cd frontend && npm run dev"
echo.
echo Kedua server sedang berjalan di jendela terminal masing-masing.
echo Silakan buka browser Anda di:
echo -> Website Utama : http://localhost:5173/
echo -> Portal CMS Admin : http://localhost:5173/admin/login
echo.
pause

