Websitemin API, self-host edilen servisler hakkında bilgileri tutuyor, bu servisler hakkında
haberleri ve güncellemeleri bir Atom feed'i aracılığı ile paylaşmama izin veriyor ve ziyartçi
metriklerini takip ediyor.

Bu dökümentasyon tüm erişime açık API endpoint'leri hakkında bilgiler içeriyor. Tüm endpoint'lere
`/api` yolu ile erişilebilir.

## Versyion 1 Endpoint'leri
Tüm versiyon 1 endpoint'leri `/v1` yolu ile erişilebilir.

Tüm endpoint'ler JSON ile formatlanmış veri döndürür.

### Hatalar
Herhangi bir hata ortaya çıkarsa 200 dışı bir cevap alırsınız. Ve JSON verisinde
bir `error` girdisi bulunur, bu ortaya çıkan hata hakkında gerekli bilgileri, metin formunda
içerir. 200 dışı bir cevap aldığınızda tek JSON girdisi bu olacaktır.

### Sonuçlar
Eğer bir hata ortaya çıkmaz ise, `error` girdisi boş bir metin olarak ayarlanır ("").
Eğer endpoint herhangi bir veri döndürüyorsa, bu veri `result` giridisi ile sağlanır.
Her endpoint için `result` girdisinin tipi farklı olabilir.

### Multilang
Bazı `result` formatları "Multilang" isimli bir yapıyı kullanabilir. Bu her desteklenen
dil için bir girdi bulunduran basit bir JSON yapısıdır. Her girdi ifade ettiği dil
ile isimlendirilir. Şuan tek desteklenen diller:
- English (`en`)
- Turkish (`tr`)

Yani her multilang yapısında, bu girdilerden **en az** bir tanesi bulunur.

İşte örnek bir multilang yapısı:
```
{
  "en": "Hello, world!",
  "tr": "Merhaba, dünya!"
}
```
Bu dökümantasyonun geri kalanında, eğer bir `result` girdisi bir multilang yapısı kullanıyorsa,
"Multilang" olarak isimlendirilecek.

### Yönetici yolları
`/v1/admin` yolu altındaki endpoint'ler yöneticiye-özeldir. Bu yollara erişmek için,
`Authorization` header'ı aracılığı ile bir parola belirtmeniz gerekecektir. Eğer
belritiğiniz parola `API_PASSWORD` ortam değişkeni ile belirtilen parola ile
uyuşuyorsa, yola erişebilirsiniz.

### GET /v1/services
Erişilebilir servislerin bir listesini döndürür. Her servis şu JSON formatını
takip eder:
```
{
  "name": "Test Service",
  "desc": {
    "en": "Service used for testing the API",
    "tr": "API'ı test etmek için kullanılan servis"
  },
  "check_time": 1735861944,
  "check_res": 1,
  "check_url": "http://localhost:7001",
  "clear": "http://localhost:7001",
  "onion": "",
  "i2p": ""
}
```
Burada:
- `name`: Servis ismi (metin)
- `desc`: Servis açıklaması (Multilang)
- `check_time`: Servisin en son durumunun kontrol edildiği zaman, eğer bu servis için
durum kontrolü desteklenmiyorsa/durum kontrolü devra dışı bırakılmış ise 0 olarak
ayarlanır (sayı, UNIX zaman damgası)
- `check_res`: En son servis durum kontrolünün sonucu (sayı)
  * servis kapalı ise 0
  * servis çalışıyor ise 1
  * serivs çalışıyor, ama yavaş ise 2
  * bu servis için durum kontrolü desteklenmiyorsa/durum kontrolü devre dışı ise 3
- `check_url`: Servis durum kontrolü için kullanılan URL (metin, yoksa boş metin)
- `clear`: Servisin açık ağ URL'si (metin, yoksa boş metin)
- `onion`: Servisin Onion (TOR) URL'si (metin, yoksa boş metin)
- `i2p`: Servisin I2P URL'si (metin, yoksa boş metin)

`name` isimli bir URL sorgusu ile servisin ismini belirterek, spesifik bir servis hakkında
bilgi de alabilirsiniz.

### GET /v1/news/:language
Verilen dil için haberlerin bir Atom feed'i döndürür. Multilang tarafından desteklenen
dilleri destekler.

### GET /v1/metrics
API kullanımı hakkınadaki metrikleri döndürür. Metrik şu formatı kullanır:
```
{
  "since":1736294400,
  "total":8
}
```
Burada:
- `since`: Metrik toplama başlangıç tarihi (sayı, UNIX zaman damgası)
- `total`: Toplam ziyaretçi sayısı (sayı)

### GET /v1/admin/logs
Yönetici kayıtlarının bir listesini döndürür. Her kayıt şu JSON formatını takip eder:
```
{
  "action": "Added service \"Test Service\"",
  "time": 1735861794
}
```
Burada:
- `action`: Yöneticinin yaptığı eylem (metin)
- `time`: Yönetici eylemin yapıldığı zaman (sayı, UNIX zaman damgası)

### PUT /v1/admin/service/add
Yeni bir servis oluşturur. İstek gövdesinin servis için kullanılan JSON formatını
takip eden JSON verisini içermesi gerekir. Bu formatı görmek için `/v1/services/all`
yoluna bakınız.

Başarılı ise herhangi bir veri döndürmez.

### DELETE /v1/admin/service/del
Bir servisi siler. İstemcinin `name` URL sorgusu ile silinecek servisin ismini belirtmesi
gerekir.

Başarılı ise herhangi bir veri döndürmez.

### GET /v1/admin/service/check
Tüm servisler için bir durum kontrolünü zorlar.

Başarılı ise herhangi bir veri döndürmez.

### PUT /v1/admin/news/add
Yeni bir haber paylaşımı oluşturur. İstek gövedisinin JSOn verisi içermesi ve verilen formatı
takip etmesi gerekir:
```
{
  "id": "test_news",
  "title": {
    "en": "Very important news",
    "tr": "Çok önemli haber"
  },
  "author": "ngn",
  "content": {
    "en": "Just letting you know that I'm testing the API",
    "tr": "Sadece API'ı test ettiğimi bilmenizi istedim"
  }
}
```
Burada:
- `id`: Haber paylaşımının benzersiz ID'si (metin)
- `title`: Haber paylaşımının başlığı (Multilang)
- `author`: Haber paylaşımının yazarı (metin)
- `content`: Haber paylaşımının içerği (Multilang)

Başarılı ise herhangi bir veri döndürmez.

### DELETE /v1/admin/news/del
Bir haber paylaşımı siler. İstemcinin `id` URL sorgusu ile silinecek paylaşımın ID'sini
belirtmesi gerekir.

Başarılı ise herhangi bir veri döndürmez.
