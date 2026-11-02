module.exports = {
  apps: [
    {
      name: 'alhamid-backend',
      cwd: '/var/www/alhamidcintamulya.com/backend',
      script: './server',
      exec_mode: 'fork',
      instances: 1,
      autorestart: true,
      watch: false,
      max_memory_restart: '500M',
      env: {
        NODE_ENV: 'production',
        PORT: 8088
      },
      error_file: '/var/www/alhamidcintamulya.com/logs/backend-error.log',
      out_file: '/var/www/alhamidcintamulya.com/logs/backend-out.log',
      time: true
    },
    {
      name: 'alhamid-frontend',
      cwd: '/var/www/alhamidcintamulya.com/frontend',
      script: 'build/index.js',
      exec_mode: 'cluster',
      instances: 'max',
      autorestart: true,
      watch: false,
      max_memory_restart: '1G',
      env: {
        NODE_ENV: 'production',
        PORT: 3005,
        HOST: '127.0.0.1',
        ORIGIN: 'https://alhamidcintamulya.com'
      },
      error_file: '/var/www/alhamidcintamulya.com/logs/frontend-error.log',
      out_file: '/var/www/alhamidcintamulya.com/logs/frontend-out.log',
      time: true
    }
  ]
};

