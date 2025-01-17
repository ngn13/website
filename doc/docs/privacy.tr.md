Kişisel olarak ben bir gizlik savunucusu olduğumdan, bir yandan sunucumu güvende tutarken bir yandan da gizliliğinize önem göstermek
için elimden geleni yapıyorum. Aynı zamanda bu tarz şeyler hakkında şeffaf ve açık olmanın önemli olduğunu düşünüyorum, o yüzden
verilerinizi nasıl işlediğimi ya da depoladığımı anlamanız için bu dökümanı yazmaya karar verdim.

## DNS & SSL
Şuan cloudflare'in isim sunucularını kullanıyorum, ancak cloudflare alan adıma sahip değil (alan adımı cloudflare'den almadım)
ve aynı şekilde herhangi bir trafiğe vekillik etmiyor. Tüm DNS kayıtlarım *Sadece DNS* modunu kullanıyor, yani sadece
DNS kayıtlarından ibaretler ve benim sunucuma işaret ediyorlar, cloudflare'e değil. Bu aynı zamanda cloudflare SSL sertifikalarımı
kontrol etmiyor demek. Tüm sertifikalar benim sunucumda tutuluyor ve Let's Encrypt ile oluşturuldular. Yani sertifikalar bana ait
ve cloudflare'in aniden DNS kayıtlarını değiştirmesi mümkün değil (bu SSL'in bozulmasına sebep olur).

## Kullanım metrikleri
Sunucumda herhangi bir istek ya da trafik monitörlermesi yok. Yani hayır, HTTP(S) istekleriniz ya da diğer ağ
bağlantılarınız renki grafikler, pasta grafikleri gibi şeyler üretmek için işlenmiyor.

Bu sayfanın altında bir ziyaretçi sayısı takipçisi olduğunu farketmiş olabilirsiniz. Bu kullandığım tek kullanım/ziyaretçi
metrik takibi ve websitemin, özgür olan, bu yüzden kodunu kendiniz denetleyebileceğiniz API'ı ile implemente edildi.

Bu metrik takipçisinin, HTTP(S) istekleriniz hakkında herhangi bir veriyi bir veri tabanına kaydetmediğini belirtmek isterim.
Bu takipçi geçici olarak IP adresinizin SHA1 hash'ini bellekte tutuyor, bunun amacı aynı ziyaretçiyi sayfayı yenilediği zaman
ya da kısa bir süre için websitesini birden fazla kez ziyaret ettiği zaman tekrar saymayı önlemek. Belirli bir miktar istekten
sonra, IP adresinizin SHA1 hash'i bellekten kaldırılacaktır ve yeni bir ziyaretçinin SHA1'i onun yerine geçicektir.

## Kayıtlar
Tüm HTTP(S) servisleri nginx ile vekilleniyor, ve nginx hepsini disk üzerindeki bir dosyaya kaydediyor. Bu dosya (`access.log`)
sadece root kullanıcısı tarafından okunabilir, ve içerği her 4 saatde bir siliniyor (diskde veri kalmadığından emin olmak için
shred komutu ile). Kayıtlar *sadece* aşağıdaki bilgileri içeriyor:

- İstek zamanı
- İstenilen host
- İstenilen yol
- HTTP istek yöntemi
- HTTP cevap kodu

Bu birşeyler yanlış giderse sorunları bulmak için ihtiyacım olan en az bilgi, kayıt tutmamın ana sebeplerinden bir tanesi
zaten bu, sorunları bulmayı kolaylaştırmak.

## Veri silimi
Sunucumdan herhangi bir verinizi kaldırmak isterseniz, [bana bir email gönderebilirsiniz](mailto:ngn@ngn.tf). Ve evet buna
kullanım metrikleri ve kayıtlar dahil.
