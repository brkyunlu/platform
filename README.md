# Platform Projesi

Bu proje, bir platformun temel işlevlerini gerçekleştirmek üzere tasarlanmıştır. Ürünlerin yönetimi, kampanyaların oluşturulması ve izlenmesi, siparişlerin oluşturulması gibi işlemleri içerir.

## Proje Yapısı

Proje dosyalarının organizasyonu şu şekildedir:

- `main.go`: Ana uygulama giriş noktası. Veritabanı bağlantıları ve temel ayarlamalar burada yapılır.
- `cmd/`: Komut işlemlerinin bulunduğu paket.
    - `command.go`: Temel komut.
    - `product.go`: Ürün işlemleri komutları.
    - `campaign.go`: Kampanya işlemleri komutları.
    - `order.go`: Sipariş işlemleri komutları.
- `controllers/`: Komut işlemlerinin yürütüldüğü paket.
    - `handlers`: Komutların init edildiği yer.
- `internal/models`: Veritabanı modelleri ve veritabanı işlemleri bulunan paket.
    - `command.go`: Komut modeli ve işlemleri.
    - `product.go`: Ürün modeli ve işlemleri.
    - `campaign.go`: Kampanya modeli ve işlemleri.
    - `order.go`: Sipariş modeli ve işlemleri.


## Kurulum

Bu projeyi çalıştırmak için aşağıdaki adımları takip edebilirsiniz:

1. Terminali açın ve projenin kök dizinine gidin.
2. `docker compose up` komutunu çalıştırarak uygulamayı başlatın.


## Kullanım

1. `docker ps` komutunu kullanarak platform-app container-id'yi kopyalayın.
2. `docker exec -it <container-id> sh` komutuyla projenin terminalini açın.
3. `./scenario.sh` komutunu yazarak senaryo dosyasını çalıştırıp komutların nasıl çalıştığını gözlemleyebilirsiniz.

## Komutlar

Veya aşağıdaki komutları tek tek oluşturabilirsiniz.

- `./main get_product_info [product_code]`
- `./main create_product [product_code] [product_price] [product_stock]`
- `./main create_order [product_code] [quantity]`
- `./main increase_time [time]` 
- `./main get_campaign_info [campaign_name]`
- `./main create_campaign [campaign_name] [product_code] [campaign_duration] [campaign_price_limit] [campaign_target_sales]`


## Proje hakkında

- Projede komutlar `cobra` kütüphanesi kullanılarak oluşturuldu.
- Veritabanı olarak PostgreSQL kullanıldı.
- Veritabanındaki verileri kontrol etmek isterseniz pgAdmin kurdum. Kullanıcı adı ve şifresine env dosyasından ulaşabilirsiniz.
