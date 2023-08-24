# Platform Projesi

Bu proje, bir platformun temel işlevlerini gerçekleştirmek üzere tasarlanmıştır. Ürünlerin yönetimi, kampanyaların oluşturulması ve izlenmesi, siparişlerin oluşturulması gibi işlemleri içerir.

## Proje Yapısı

Proje dosyalarının organizasyonu şu şekildedir:

- `main.go`: Ana uygulama giriş noktası. Veritabanı bağlantıları ve temel ayarlamalar burada yapılır.
- `cmd/commands`: Komut işlemlerinin bulunduğu paket.
    - `command.go`: Temel komut ve komut yönetimi arayüzü.
    - `product.go`: Ürün işlemleri komutları.
    - `campaign.go`: Kampanya işlemleri komutları.
    - `order.go`: Sipariş işlemleri komutları.
    - `commandprocessor.go`: Kullanıcıdan komutları alıp işleyen komut işleyici.
    - `input.go`: Kullanıcı giriş işlemleri yardımcı fonksiyonları.
- `internal/models`: Veritabanı modelleri ve veritabanı işlemleri bulunan paket.
    - `command.go`: Komut modeli ve işlemleri.
    - `product.go`: Ürün modeli ve işlemleri.
    - `campaign.go`: Kampanya modeli ve işlemleri.
    - `order.go`: Sipariş modeli ve işlemleri.
- `README.md`: Proje hakkında detaylı açıklamaların yer aldığı belge.

## Kurulum

Bu projeyi çalıştırmak için aşağıdaki adımları takip edebilirsiniz:

1. `.env` dosyasını projenin kök dizinine ekleyin ve gerekli ortam değişkenlerini ayarlayın.
2. Terminali açın ve projenin kök dizinine gidin.
3. `go run main.go` komutunu çalıştırarak uygulamayı başlatın.
4. Komut işleyicisi çalıştığında, komutları izleyerek ürün ve kampanya işlemlerini gerçekleştirebilirsiniz.

## Kullanım

Aşağıda projede kullanılabilir komutlar ve nasıl kullanılacaklarına dair örnekler bulunmaktadır.

### Ürün İşlemleri

- **get_product_info**: Bir ürünün bilgilerini alır.
    - Kullanım: `get_product_info -ÜrünKodu`

- **create_product**: Yeni bir ürün oluşturur.
    - Kullanım: `create_product -ÜrünKodu -Fiyat -Stok`

### Kampanya İşlemleri

- **get_campaign_info**: Bir kampanyanın bilgilerini alır.
    - Kullanım: `get_campaign_info -KampanyaAdı`

- **create_campaign**: Yeni bir kampanya oluşturur.
    - Kullanım: `create_campaign -KampanyaAdı -ÜrünKodu -Süre -FiyatManipülasyonLimiti -HedefSatış`

### Sipariş İşlemleri

- **create_order**: Bir sipariş oluşturur.
    - Kullanım: `create_order -ÜrünKodu -Miktar`

### Zaman İşlemleri

- **increase_time**: Zamanı belirtilen saat kadar ilerletir.
    - Kullanım: `increase_time -Saat`

### Çıkış

- **exit**: Uygulamadan çıkış yapar.
    - Kullanım: `exit`

## Kullanım Örnekleri

- Ürün bilgilerini almak için: `get_product_info ABC123`
- Yeni bir ürün oluşturmak için: `create_product XYZ456 29.99 50`
- Kampanya bilgilerini almak için: `get_campaign_info YazKampanyası`
- Yeni bir kampanya oluşturmak için: `create_campaign YazKampanyası ABC123 30 20.0 100`
- Sipariş oluşturmak için: `create_order ABC123 5`
- Zamanı ilerletmek için: `increase_time 3`
- Uygulamadan çıkış yapmak için: `exit`

Bu komutlar sayesinde ürünleri yönetebilir, kampanyalar oluşturabilir, siparişler oluşturabilir ve zamanı ilerletebilirsiniz. Her bir komutun ardından gelen açıklamaları ve örnek kullanımları takip ederek işlemlerinizi gerçekleştirebilirsiniz.

---

Projede yer alan kodlar ve komutlar, farklı işlemleri gerçekleştirmek üzere tasarlanmıştır. Daha fazla detay ve kullanım için ilgili kod dosyalarını inceleyebilirsiniz.

## Geliştirme

Projenin geliştirilmesi ve genişletilmesi için aşağıdaki adımları izleyebilirsiniz:

1. İlgili komut dosyalarını (`product.go`, `campaign.go`, vb.) düzenleyerek yeni komutlar ekleyebilirsiniz.
2. Veritabanı modellerini (`models` altındaki dosyalar) gerektiği gibi güncelleyebilirsiniz.
3. Yeni özellikler eklemek için gerekli işlemleri `cmd/commands` altındaki uygun dosyalara ekleyebilirsiniz.

## Lisans

Bu proje MIT lisansı altında lisanslanmıştır. Daha fazla bilgi için [LICENSE](LICENSE) dosyasını inceleyebilirsiniz.
