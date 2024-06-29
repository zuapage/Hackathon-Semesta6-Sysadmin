# Hackathon-Semesta6-Sysadmin
Repositori berikut merupakan solusi saya dalam memecahkan masalah yang telah diberikan dalam rangka memenuhi seleksi hackaton beasiswa semesta untuk posisi System Administrator.

# Automation and Secure Deployment
Lab ini dibangun menggunakan virtualisasi VirtualBox, dengan dua VM yang berperan sebagai server development dan server production. Beberapa alat digunakan untuk menunjang proses pemecahan masalah dalam kasus ini.

![semesta-hackaton](https://github.com/zuapage/Hackathon-Semesta6-Sysadmin/assets/68767230/99cadd82-46ef-498a-b18e-3da0d0ddd6e6)

## -- Setup & Installation --
1. Menjalankan Script Preparation
   Jalankan `preparation.sh` yang berada pada direktori `preparation`. Script ini akan melakukan:
   - Instalasi Docker dan net-tools.
   - Membuat jaringan Docker dengan nama `dev_semesta` yang akan digunakan oleh kontainer (Ansible, Jenkins, dan Grafana).
2. Menjalankan Ansible Menggunakan Docker Compose
   Ansible akan dijalankan menggunakan Docker Compose untuk manajemen yang lebih mudah dan terisolasi.
3. Menjalankan Ansible Playbook
   - Menggunakan roles untuk konfigurasi yang lebih terstruktur pada server `dev-semesta` dan `prod-semesta`.
   - Penggunaan non-default SSH port pada inventory yang berjalan di port 1566 untuk keamanan tambahan.
4. Setup Tunneling dengan Cloudflare
   Menyiapkan tunneling pada Cloudflare untuk mengekspos layanan Jenkins untuk webhook dan aplikasi Semesta yang berjalan pada server `dev-semesta` dan `prod-semesta`.
5. Membuat dan Menjalankan Jenkins Pipeline
   Membuat pipeline untuk workflow CI/CD yang mencakup:
   - Programmer melakukan push ke repository GitHub.
   - Trigger job Jenkins secara otomatis menggunakan webhook dengan tunneling pada Cloudflare.
   - Tahapan pipeline meliputi:
     - Clone repository dari SCM.
     - Menghapus image build sebelumnya.
     - Build aplikasi `semesta-app1` dan `semesta-app2` dengan Dockerfile secara paralel.
     - Push image ke Docker registry.
     - Deploy aplikasi `semesta-app1` dan `semesta-app2` di server `prod-semesta`.
6. Setup Nginx (Reverse Proxy)
   Mengonfigurasi Nginx sebagai reverse proxy untuk mengarahkan permintaan ke aplikasi yang berjalan pada container Docker.
7. Setup Monitoring
   Menggunakan Telegraf sebagai agen pengumpul metrik, Prometheus untuk penyimpanan metrik, dan Grafana untuk visualisasi data. Monitoring dilakukan untuk Docker dan Nginx.

## -- CI/CD Workflow --
- Programmer melakukan push ke repository GitHub menggunakan Git.
- Jenkins secara otomatis menjalankan job saat terjadi push dengan webhook yang diatur melalui Cloudflare.
- Job Jenkins memiliki beberapa tahap atau stage:
  - Clone repository dari SCM.
  - Menghapus image build sebelumnya.
  - Build aplikasi `semesta-app1` dan `semesta-app2` dengan Dockerfile secara paralel.
  - Push image ke Docker registry.
  - Deploy aplikasi `semesta-app1` dan `semesta-app2` di server `prod-semesta`.

Berikut adalah tampilan trigger build dengan webhook pada pipeline job:

<img width="800" alt="Screenshot 2024-06-30 023807" src="https://github.com/zuapage/Hackathon-Semesta6-Sysadmin/assets/68767230/6a46eb67-4723-421e-ade5-c74a204cef35">

Berikut adalah tampilan ketika pipeline job berhasil dijalankan:

<img width="800" alt="Screenshot 2024-06-29 223341" src="https://github.com/zuapage/Hackathon-Semesta6-Sysadmin/assets/68767230/63865525-0e5c-4349-8b38-367b0fc2b177">

Berikut adalah tampilan dari aplikasi yang telah berhasil di deploy dari hasil pipeline job ke server 'prod-semesta':

<img width="900" alt="Screenshot 2024-06-29 223440" src="https://github.com/zuapage/Hackathon-Semesta6-Sysadmin/assets/68767230/2bf4b5b3-5bc1-4efa-b004-d87bedc8e897">


## -- Monitoring --
Setelah aplikasi di-deploy ke server production, monitoring dilakukan untuk memantau kinerja aplikasi. Berikut adalah contoh tampilan monitoring menggunakan Grafana:

<img width="959" alt="DASHBOARDD" src="https://github.com/zuapage/Hackathon-Semesta6-Sysadmin/assets/68767230/22e6df67-8951-4512-8ec6-c30d7cc4cafd">


## -- User Workflow --
- User mengakses domain `semesta.mafumaku.biz.id`.
- Permintaan akan diteruskan melalui reverse proxy ke port 3000 tempat `semesta-app1` berjalan.
- Berikut adalah respon yang diterima user ketika permintaan berhasil diproses:

<img width="959" alt="semesta-app1" src="https://github.com/zuapage/Hackathon-Semesta6-Sysadmin/assets/68767230/585194eb-313e-4a15-93ab-3b8ead070727">


- Jika user mengakses domain `semesta.mafumaku.biz.id` dengan endpoint `/aboutus`:
  - Permintaan akan diteruskan melalui reverse proxy ke port 3001 tempat `semesta-app2` berjalan.
  - Berikut adalah respon yang diterima user ketika permintaan berhasil diproses:

<img width="959" alt="semesta-app2" src="https://github.com/zuapage/Hackathon-Semesta6-Sysadmin/assets/68767230/7f844d68-1fa2-48fc-917d-6144a20ece81">


