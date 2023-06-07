package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"backend/api/models"
)

var posts = []models.Post{

	//==published==//	
	models.Post{
		Title:   "Minyak Ikan Kod dan Minyak Ikan Tuna, Mana yang Lebih Baik?",
		Content: "Nutrisi memiliki peran penting dalam mendukung proses tumbuh kembang si Kecil. Di antara berbagai asupan nutrisi yang penting untuk anak, Bunda jangan sampai melewatkan asupan minyak ikan, ya! Karena, manfaat minyak ikan untuk anak sangatlah banyak, terutama bagi proses perkembangan otak yang optimal.",
		Category: "Nutrisi dan Makanan",
		Status: "publish",
	},
	models.Post{
		Title:   "Ketahui Apa Saja Hak dan Kewajiban Anak di Rumah",
		Content: "Setiap anak berhak mendapatkan pendidikan. Tidak cuma dalam bentuk pendidikan formal dengan cara menyekolahkan anak, tapi juga mendidik anak belajar di rumah. Misalnya, ketika si Kecil merasa kesulitan dengan tugas sekolahnya, Bunda harus membantu mengajari anak sampai ia mengerti. Tidak hanya pelajaran sekolah saja, Bunda juga bisa mengajarkan anak soal pelajaran hidup, seperti tentang moral, tata krama, dan etika berinteraksi dengan orang lain supaya si Kecil nantinya tumbuh menjadi pribadi yang lebih baik.",
		Category: "Nutrisi dan Makanan",
		Status: "publish",
	},
	models.Post{
		Title:   "Cara Sederhana Meringankan Batuk Pilek pada Anak",
		Content: "Batuk pilek umumnya disebabkan oleh virus atau bakteri yang menginfeksi hidung dan tenggorokan. Batuk pilek karena infeksi virus dapat sembuh sendiri dalam waktu beberapa hari, sehingga tidak perlu diobati dengan antibiotik. Yang perlu diobati dengan antibiotik adalah batuk pilek karena infeksi bakteri. Meski umumnya tidak berbahaya, batuk pilek dapat mengganggu aktivitas dan istirahat anak. Si Kecil yang biasanya ceria juga bisa jadi terlihat lesu dan tidak bersemangat. Bila seperti ini, Anda sebagai orang tua pasti merasa sedih dan ingin Si Kecil bisa segera kembali sehat. Jika Si Kecil menderita batuk pilek, Anda disarankan untuk berkonsultasi ke dokter terlebih dahulu, terutama apabila Si Kecil belum berusia 2 tahun.",
		Category: "Keluarga",
		Status: "publish",
	},
	models.Post{
		Title:   "Rekomendasi Balsem Bayi yang Aman dan Nyaman untuk Kulit Bayi",
		Content: "Balsem bayi dapat digunakan untuk membantu melegakan pernapasan saat Si Buah Hati batuk atau pilek. Balsem bayi juga bisa memberikan rasa hangat pada badan bayi dan membantu meredakan perut kembung. Perut kembung, batuk, dan pilek merupakan kondisi yang sering dialami oleh bayi. Saat mengalami keluhan tersebut, ia akan rewel akibat sulit bernapas, sakit kepala, hidung tersumbat, perut terasa penuh, atau karena badannya terasa tidak nyaman.",
		Category: "Kesehatan",
		Status: "publish",
	},
	models.Post{
		Title:   "Beragam Obat Diare Bayi yang Aman",
		Content: "Obat diare bayi dapat diberikan untuk mencegah terjadinya komplikasi, seperti dehidrasi. Ada beragam obat diare yang aman diberikan untuk bayi, mulai dari memberikan ASI lebih sering hingga probiotik. Bayi dikatakan mengalami diare jika ada perubahan pada tekstur, frekuensi, dan warna tinjanya. Diare pada bayi bisa disebabkan oleh banyak hal, seperti infeksi, perubahan tekstur makanan dari cair ke padat, atau kondisi medis tertentu.",
		Category: "Keluarga",
		Status: "publish",
	},

	//==draft==//	
	models.Post{
		Title:   "6 Warna Feses Bayi Diare",
		Content: "Memperhatikan warna feses bayi diare dapat menjadi cara untuk mengetahui kondisi bayi. Berdasarkan warna feses bayi, penyebab diare pada bayi dapat dibedakan. Selain itu, mengetahui warna feses bayi diare pun dapat menjadi petunjuk kapan harus membawanya ke dokter untuk mendapatkan penanganan.",
		Category: "Keluarga",
		Status: "draft",
	},
	models.Post{
		Title:   "Inilah 9 Perlengkapan Bayi yang Tidak Perlu Dibeli",
		Content: "Perlengkapan bayi yang lucu sering kali membuat Bumil lapar mata. Namun, ada beragam perlengkapan bayi yang tidak perlu dibeli, lho. Selain bisa membuat Bumil lebih berhemat, tetapi juga menghindari kebiasaan menumpuk barang.",
		Category: "Keluarga",
		Status: "draft",
	},
	models.Post{
		Title:   "Agar Aman dan Nyaman, Ini Tips Membawa Bayi Saat Perjalanan Mudik",
		Content: "Membawa bayi dalam perjalanan mudik Lebaran bisa saja terasa merepotkan, apalagi jika ini adalah mudik pertama Bunda dengan Si Kecil. Namun, tidak perlu terlalu khawatir ya, Bun. Jika dipersiapkan dengan baik, perjalanan mudik dengan Si Kecil tetap bisa terasa nyaman dan menyenangkan, kok.",
		Category: "Keluarga",
		Status: "draft",
	},
	models.Post{
		Title:   "Bayi Jarang Menangis, Apakah Hal Ini Normal?",
		Content: "Menangis adalah cara bayi untuk berkomunikasi dengan Anda. Alasan bayi menangis biasanya untuk memberi tahu bahwa dia lapar, mengantuk, kedinginan atau kepanasan, popoknya basah, tidak nyaman, takut, atau bosan dan sekadar ingin digendong.",
		Category: "Keluarga",
		Status: "draft",
	},

	//==trashed==//	
	models.Post{
		Title:   "3 Kriteria Tempat Tidur Bayi yang Ideal",
		Content: "Memilih tempat tidur bayi bukan perkara yang mudah karena ada beberapa kriteria yang perlu dipenuhi. Ini penting untuk diperhatikan, sebab pemilihan tempat tidur yang tepat tidak hanya membuat Si Kecil tertidur nyaman, tapi juga terhindar dari bahaya. ",
		Category: "Keluarga",
		Status: "trash",
	},
	models.Post{
		Title:   "Amankah Posisi Bayi Tidur Miring?",
		Content: "Pada orang dewasa, posisi tidur miring memang dapat menyebabkan ngiler saat tidur, tetapi posisi ini bermanfaat untuk mencegah mendengkur, mengurangi mulas, menjaga kesehatan saluran cerna, dan menghindari nyeri punggung bawah. Namun, tidur miring bukanlah posisi tidur yang disarankan untuk bayi. Pasalnya, bayi yang tidur miring sering kali berakhir dengan posisi tidur tengkurap. Posisi tidur miring dan tengkurap bisa saja membuat saluran pernapasan bayi menyempit atau terhalang. Akibatnya, bayi menjadi kesulitan bernapas dan kekurangan oksigen. Kondisi ini juga akan meningkatkan risiko sindrom kematian bayi mendadak (SIDS).",
		Category: "Keluarga",
		Status: "trash",
	},
	models.Post{
		Title:   "Daftar Makanan Bergizi untuk Ibu Hamil dan Manfaatnya",
		Content: "Makanan bergizi untuk ibu hamil sangat berpengaruh terhadap kesehatan ibu dan pertumbuhaan janin. Oleh karena itu, penting bagi ibu hamil untuk mengetahui pilihan makanan bergizi yang bisa dikonsumsi agar janin yang dikandung tetap sehat. Konsumsi makanan bergizi untuk ibu hamil dilakukan agar kebutuhan nutrisi saat hamil tercukupi dengan baik. Meski demikian, ibu hamil perlu untuk memerhatikan apa yang dikonsumsinya. Pasalnya, ada makanan bergizi yang tidak diperbolehkan untuk dikonsumsi selama kehamilan, seperti telur dan daging mentah, serta produk olahan susu yang belum melewati proses pasteurisasi.",
		Category: "Keluarga",
		Status: "trash",
	},
	models.Post{
		Title:   "Kapan Si Kecil Boleh Mengonsumsi Yoghurt?",
		Content: "Yoghurt dikenal baik untuk kesehatan, terutama dalam memelihara kesehatan saluran pencernaan. Namun, apakah yoghurt untuk bayi boleh diberikan sebagai MPASI? Mulai usia berapakah yoghurt sebaiknya diberikan kepada bayi? Simak jawabannya dalam artikel berikut. Yoghurt merupakan produk olahan susu hasil fermentasi. Selain mengandung nutrisi yang sama dengan susu, yoghurt juga mengandung probiotik. Probiotik bermanfaat dalam meningkatkan sistem kekebalan tubuh dan kesehatan saluran cerna.",
		Category: "Keluarga",
		Status: "trash",
	},
	models.Post{
		Title:   "Bayi Ngorok, Kenali Penyebab dan Cara Mencegahnya",
		Content: "Bayi ngorok di minggu awal kelahirannya adalah kondisi yang normal terjadi. Namun, Bunda harus tetap waspada, karena ini bisa menjadi tanda Si Kecil mengalami gangguan kesehatan bila terjadi secara terus-menerus atau suara ngorok terdengar sangat keras. Bayi baru lahir yang sedang tidur biasanya bernapas sambil mengeluarkan suara atau dengkuran. Hal ini karena saluran pernapasannya masih sempit dan berisi banyak lendir. Suara ngorok pada bayi biasanya hilang ketika saluran pernapasannya telah berkembang sempurna dan ia sudah mampu menelan ludah.",
		Category: "Keluarga",
		Status: "trash",
	},
	models.Post{
		Title:   "Kenali Penyebab Bayi Kuning dan Cara Mengatasinya",
		Content: "Bayi kuning atau penyakit kuning merupakan kondisi yang umum dialami bayi baru lahir. Meski umumnya tidak berbahaya, ada beberapa hal yang perlu diperhatikan bila bayi menunjukkan tanda-tanda penyakit kuning. Dengan begitu, dapat segera dilakukan penanganan yang tepat. Kondisi bayi kuning ditandai dengan warna kuning pada kulit atau bagian putih mata bayi. Selain itu, bayi yang mengalami penyakit kuning biasanya memiliki urine yang berwarna kuning pekat, tinja berwarna pucat, serta telapak tangan dan kakinya pun menguning.",
		Category: "Nutrisi dan Makanan",
		Status: "trash",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	

	for i, _ := range posts {
		
		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}