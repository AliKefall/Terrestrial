Bu proje dönem ödevi olarak hazırlanmıştır aşağıda bu proje ile ilgili dökümantasyonu bulabilirsiniz.



Teknik Tercihler:
1- Backenddeki api yapısı için golang dili kullanılmıştır. Hem daha fazla hakim olmamdan dolayı
hem de routerları local ortamda daha kolay kurulup test edilebildiği için seçtim. Bunun yanındasqlc gibi crud işlemlerini benim için fonksiyona dökebilen yapılar(./internal/db/) olduğu için tercih edilmiştir.
2- Router olarak chi paketi kullanılmıştır. Özellikle düşük ölçekte çok fazla konfigurasyon ayarlaması istemediği için hem de loglama tekniklerinin basitliğinden dolayı seçilmiştir.
3- Session token olarak da bilinen bir kişinin her yaptığı request için authentication karşılığı alma zorunluluğunu ortadan kaldıran JWT tokenlar kullanılmıştır .env dosyasındaki TOKEN_SECRET değişken adı ise JWT şifrelemesi için gerekli olan rastgele bir değerdir ama siz istediğiniz değer ile değiştirebilirsiniz.
4- Database olarak Turso kullanılmıştır. Hem golang için sağlam bir kütüphanesi olması hem de ücretsiz versiyonunun bile gerekenden fazla trafik hakkı vermesinden dolayı kullanılmıştır. Turso bağlantısı için gereken database connection key ve database authentication key .env dosyası içerisine yazılmalıdır. 
5- Kullanıcı şifrelerini SHA256 ile şifreledim. Projenin en başındayken IV ve salt keyler ile kendi şifreleme sistemimi de yapmıştım ancak bu ölçek için gereksiz derecede işleri karmaşıklaştırıp fazladan teknik anlatım zorunluluğu çıkarırdı. Bycrypt kütüphanesi şu hali ile bile gayet yeterlidir.
6- Sqlc yapısı ise bir çok dilde olduğu gibi query yapısını fonksiyon haline getirmemize yarayan bir extension pack olarak görev görür. Yaptığı en yegane şey CRUD işlemleri için rahatlık sağlamasıdır. Ancak sonradan da belirtileceği üzere kullandığımız databasein kullandığı sqlite3 driveri yüzünden özellikle de datetime yapılarını doğru çıkarabilmesi için sqlc.yaml konfigürasyon dosyasında hem uuid yapısı için hem de datetime yapıları için overridelar yazılmıştır.

Projenin Amacı:
Bu proje dönem ödevi hazırlanmaktan ziyade sonradan yapılabilecek değişimler için daha doğrusu scale oldukça da mimarisi dolayısıyla geliştirmeye uygun bir yapıdır. Bu projede amacımız olan gelir-gider günlük haftalık ve yıllık dökümantasyon çıkartma işlemleri yapılmıştır. GUI yapısı ise Python ve Qt frameworku ile hazırlanmış olup backend işlemleri harici kullanıcıya görsel olarak sunmak üzere de geliştirme ortamında işlenmiştir.

Projedeki darboğazlar:
Projenin şu anki aşaması için pekte değiştirilmesi ya da refactor edilmesi gerekilen bir durum yoktur ancak daha büyük bir ölçek veya daha büyük bir feature için kesinlikle sqlc yapısı terk edilip custom queryler ile devam edilmelidir. Bunun ise en büyük sebebi sqlc go için queryleri işlerken zaman objelerini time.Time olarak işler ancak Turso libsql tabanlı bir veritabanı olduğundan dolayı veriyi sadece primitive type olarak kabul eder. Ancak sqlc time.Time ve ayrı olarak uuid yapısını da kendi kütüphanelerine ait obje yapısı ile tutar. Bu sorunu çözmek için sqlc.yaml dosyası içerisinde tamamına override eklenmiştir. Bunun bir örneği olarak transactionların çoğu için custom query kullanılmış olup sadece ama sadece crud işlemleri için sqlc yapısı kullanılmıştır. Bunun haricinde ise Router yapısı için fazladan bir concurrency yapısı yeterlidir.

Kurulum:
go versiyon 1.24.0 kullanılmalıdır.
indirmek için golang'ın kendi sitesi olan go.dev sitesini ziyaret edebilirsiniz.



Projenin dizini için:
git clone https://github.com/AliKefall/Terrestrial 

yaptıktan sonra dosya dizininde şu komutu çalıştırın.
go mod download 

Bu komut gerekli olan bütün kütüphaneleri go.mod ve go.sum dosyaları içerisinden çekerek indirir.

Proje için .env.example dosyası içerisindeki değişkenleri kendi verileriniz ile doldurmanız gerekmektedir örneğin database bağlantısı ve authentication kısmı tamamen kullanıcının sahip olup doldurması gereken yerlerdir. Bu bilgiler kesinlikle public olarak paylaşılmaz.

Bütün hepsi bittikten sonra dosyanın rootundan şu komut çalıştırılmalıdır: 

go run ./server/cmd .

Her şey doğru kurulmuşsa program localhost:8080 yapısı üzerinden çalışacaktır. Frontendin çalışması için backendin aktif olması şarttır.


