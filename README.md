
# Terrestrial

Bu proje, dönem ödevi kapsamında geliştirilmiştir. Proje; ölçeklenebilirlik göz önünde bulundurularak tasarlanmış bir backend mimarisi ve buna bağlı bir masaüstü arayüzünden oluşmaktadır.

Aşağıda proje ile ilgili teknik tercihler, amaç, mevcut sınırlamalar ve kurulum adımları detaylı şekilde açıklanmıştır.

---

## Teknik Tercihler

### Backend ve API Yapısı
Backend API yapısı için **Go (Golang)** dili kullanılmıştır. Bu tercih, dile olan hâkimiyetimin yanı sıra router yapılarını lokal ortamda daha kolay kurup test edebilme imkânı sağlaması sebebiyle yapılmıştır.

Ayrıca **sqlc** gibi SQL sorgularını Go fonksiyonlarına dönüştüren araçların (`./internal/db/`) CRUD işlemlerinde sağladığı kolaylık, Go’nun tercih edilmesinde etkili olmuştur.

---

### Router Seçimi
Router olarak **chi** paketi kullanılmıştır. Düşük ve orta ölçekli projelerde fazla konfigürasyon gerektirmemesi ve loglama mekanizmalarının sade olması nedeniyle tercih edilmiştir.

---

### Kimlik Doğrulama (Authentication)
Kullanıcıların her request için yeniden kimlik doğrulaması yapmasını engellemek amacıyla **JWT (JSON Web Token)** kullanılmıştır.

`.env` dosyasında yer alan `TOKEN_SECRET` değişkeni, JWT imzalama işlemi için kullanılan rastgele bir anahtardır ve istenildiği takdirde değiştirilebilir.

---

### Veritabanı Seçimi
Veritabanı olarak **Turso** tercih edilmiştir. Bunun başlıca sebepleri şunlardır:
- Go için stabil ve iyi dokümante edilmiş bir kütüphaneye sahip olması
- Ücretsiz planının proje için yeterli trafik limiti sunması

Turso bağlantısı için gerekli olan **database connection key** ve **authentication key** bilgileri `.env` dosyası içerisine eklenmelidir.

---

### Şifreleme
Kullanıcı şifreleri **bcrypt** algoritması kullanılarak hashlenmiştir.

Projenin ilk aşamasında IV ve salt değerleri ile özel bir şifreleme mekanizması tasarlanmış olsa da, bu ölçek için gereksiz karmaşıklık oluşturacağı ve ek teknik anlatım gerektireceği için bu yaklaşımdan vazgeçilmiştir. Bcrypt, mevcut ihtiyaçlar için yeterli güvenliği sağlamaktadır.

---

### SQLC Kullanımı
**sqlc**, SQL sorgularını Go fonksiyonlarına dönüştürmeye yarayan bir araçtır ve temel amacı CRUD işlemlerini kolaylaştırmaktır.

Kullanılan veritabanının **libsql (sqlite3 tabanlı)** olması nedeniyle bazı tip uyuşmazlıkları ortaya çıkmıştır. Özellikle:
- `time.Time`
- `uuid`

gibi tiplerin veritabanında primitive türler olarak tutulması gerekliliği sebebiyle `sqlc.yaml` dosyasında bu alanlar için override tanımlamaları yapılmıştır.

Bu nedenle:
- Basit CRUD işlemlerinde sqlc kullanılmış,
- Daha karmaşık transaction işlemlerinde custom SQL sorguları tercih edilmiştir.

---

## Projenin Amacı

Bu proje yalnızca bir dönem ödevi olarak hazırlanmakla sınırlı değildir. Aynı zamanda ileride yapılabilecek geliştirmeler ve ölçeklenme senaryoları göz önünde bulundurularak tasarlanmıştır.

Projenin temel amacı:
- Gelir ve gider kayıtlarının tutulması
- Günlük, haftalık ve yıllık bazda raporlanmasıdır

Frontend tarafı, **Python ve Qt framework** kullanılarak geliştirilmiştir. Bu yapı, backend işlemlerinden bağımsız olarak kullanıcıya görsel bir arayüz sunmak amacıyla hazırlanmıştır.

Frontend uygulamasının çalışabilmesi için backend servisinin aktif olması gerekmektedir.

---

## Mevcut Darboğazlar ve Sınırlamalar

Projenin mevcut ölçeği için acil bir refactor veya mimari değişiklik ihtiyacı bulunmamaktadır. Ancak daha büyük bir ölçek veya daha karmaşık feature’lar eklenecek olursa:

- sqlc yapısının terk edilmesi
- tamamen custom SQL sorgularına geçilmesi

daha sağlıklı olacaktır.

Bunun temel sebebi, sqlc’nin `time.Time` ve `uuid` gibi tipleri kendi Go objeleri olarak ele alması; Turso’nun ise bu verileri yalnızca primitive türler olarak kabul etmesidir. Mevcut durumda bu durum override’lar ile çözülmüştür.

Router tarafında ise daha büyük bir yük altında ek concurrency mekanizmalarına ihtiyaç duyulabilir.

---

## Kurulum

### Gereksinimler
- **Go 1.24.0**

Go kurulumu için resmi web sitesi:
- https://go.dev

---

### Projenin Klonlanması
```bash
git clone https://github.com/AliKefall/Terrestrial

-- En son olarak projeyi çalıştırmak için proje dizininde:
go run ./server/cmd . 

Komutu çalıştırılmalıdır.
