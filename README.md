# Hackathon-Semesta6-Sysadmin
Repositori berikut merupakan solusi saya dalam memecahkan masalah yang telah diberikan dalam rangka memenuhi seleksi hackaton beasiswa semesta untuk posisi System Administrator.

# Automation and Secure Deployment
Lab ini dibangun menggunakan virtualisasi VirtualBox, dengan dua VM yang berperan sebagai server development dan server production. Beberapa alat digunakan untuk menunjang proses pemecahan masalah dalam kasus ini.

![semesta-hackaton](https://github.com/zuapage/Hackathon-Semesta6-Sysadmin/assets/68767230/9dc99ba8-b59f-4531-b461-4c051420ccf8)


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

<img width="800" alt="webhook" src="https://github.com/zuapage/Hackathon-Semesta6-Sysadmin/assets/68767230/8f79392b-3958-46ee-9df4-89cecbef5a42">


Berikut adalah tampilan ketika pipeline job berhasil dijalankan:

<img width="800" alt="Screenshot 2024-06-29 223341" src="https://github.com/zuapage/Hackathon-Semesta6-Sysadmin/assets/68767230/3204d100-308b-4a75-a4af-b92add3475a8">


Berikut adalah tampilan dari aplikasi yang telah berhasil di deploy dari hasil pipeline job ke server 'prod-semesta':

<img width="800" alt="Screenshot 2024-06-29 223440" src="https://github.com/zuapage/Hackathon-Semesta6-Sysadmin/assets/68767230/d21b44d1-5325-49bf-9e22-0c7de07cdd90">


## -- Monitoring --
Setelah aplikasi di-deploy ke server production, monitoring dilakukan untuk memantau kinerja aplikasi. Berikut adalah contoh tampilan monitoring menggunakan Grafana:

<img width="959" alt="DASHBOARDD" src="https://github.com/zuapage/Hackathon-Semesta6-Sysadmin/assets/68767230/e673c9f6-7f85-42d7-b5dd-c5d30d13fe0a">


## -- User Workflow --
- User mengakses domain `semesta.mafumaku.biz.id`.
- Permintaan akan diteruskan melalui reverse proxy ke port 3000 tempat `semesta-app1` berjalan.
- Berikut adalah respon yang diterima user ketika permintaan berhasil diproses:

<img width="959" alt="semesta-app1" src="https://github.com/zuapage/Hackathon-Semesta6-Sysadmin/assets/68767230/c967a818-e905-408f-bd70-2a17dbf571b7">


- Jika user mengakses domain `semesta.mafumaku.biz.id` dengan endpoint `/aboutus`:
  - Permintaan akan diteruskan melalui reverse proxy ke port 3001 tempat `semesta-app2` berjalan.
  - Berikut adalah respon yang diterima user ketika permintaan berhasil diproses:

<img width="959" alt="semesta-app2" src="https://github.com/zuapage/Hackathon-Semesta6-Sysadmin/assets/68767230/bb96a0ce-23e6-4f6d-ab6b-77f7eda9fbb4">



