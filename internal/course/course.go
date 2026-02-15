package course

import (
	"fmt"

	"gotrack/internal/models"
)

// ExerciseRequirements maps exercise chapters to the theory chapter that must be completed first.
var ExerciseRequirements = map[int]int{
	3:  2,  // Exercicios #1 requer Cap 2
	5:  4,  // Exercicios #2 requer Cap 4
	7:  6,  // Exercicios #3 requer Cap 6
	9:  8,  // Exercicios #4 requer Cap 8
	11: 10, // Exercicios #5 requer Cap 10
	13: 12, // Exercicios #6 requer Cap 12
	15: 14, // Exercicios #7 requer Cap 14
	17: 16, // Exercicios #8 requer Cap 16
	20: 18, // Exercicios #9 requer Cap 18
}

func lesson(chapter, num int, title, ytID string, exercise bool) models.Lesson {
	return models.Lesson{
		ID:           lessonID(chapter, num),
		Chapter:      chapter,
		LessonNumber: num,
		Title:        title,
		YouTubeID:    ytID,
		IsExercise:   exercise,
	}
}

func lessonID(chapter, num int) string {
	return fmt.Sprintf("cap%02d_aula%02d", chapter, num)
}

// AllChapters returns the full course content organized by chapters.
func AllChapters() []models.Chapter {
	return []models.Chapter{
		{
			Number: 1, Title: "Visao Geral",
			Lessons: []models.Lesson{
				lesson(1, 1, "Bem-vindo!", "WiGU_ZB-u0w", false),
				lesson(1, 2, "Por que Go?", "jXA0O5b-F1g", false),
				lesson(1, 3, "Sucesso", "CvH9fC2eYO8", false),
				lesson(1, 4, "Recursos", "n7PZo55Wl5A", false),
				lesson(1, 5, "Como esse curso funciona", "V9zfX8Yhc9g", false),
				lesson(1, 6, "Update: Exercicios", "B5sy5v2dT9Y", false),
			},
		},
		{
			Number: 2, Title: "Variaveis, Valores & Tipos",
			Lessons: []models.Lesson{
				lesson(2, 1, "Go Playground", "Qf2645sTqH0", false),
				lesson(2, 2, "Hello world!", "tdZ2I2RZ7JI", false),
				lesson(2, 3, "Operador curto de declaracao", "QT7YvrEWMGE", false),
				lesson(2, 4, "A palavra-chave var", "YAGKHSNWdEE", false),
				lesson(2, 5, "Explorando tipos", "DazALie4u5E", false),
				lesson(2, 6, "Valor zero", "Ruwp2xLD_AI", false),
				lesson(2, 7, "O pacote fmt", "IjDiONSi-tI", false),
				lesson(2, 8, "Criando seu proprio tipo", "DGHYibXY6Qk", false),
				lesson(2, 9, "Conversao, nao coercao", "s0UYWJUz30k", false),
			},
		},
		{
			Number: 3, Title: "Exercicios: Nivel #1", IsExercise: true, RequiredBy: 2,
			Lessons: []models.Lesson{
				lesson(3, 1, "Exercicio 1", "5K17jFDXvWw", true),
				lesson(3, 2, "Exercicio 2", "Q7ZrIDMj9zc", true),
				lesson(3, 3, "Exercicio 3", "VkuZ4QMnoZM", true),
				lesson(3, 4, "Exercicio 4", "5-0S-gefNe0", true),
				lesson(3, 5, "Exercicio 5", "O0r318FN_Uw", true),
				lesson(3, 6, "Exercicio 6", "OINd35-02xo", true),
			},
		},
		{
			Number: 4, Title: "Fundamentos da Programacao",
			Lessons: []models.Lesson{
				lesson(4, 1, "Tipo booleano", "voisg61hPXA", false),
				lesson(4, 2, "Como os computadores funcionam", "C55_dhNmvrQ", false),
				lesson(4, 3, "Tipos numericos", "3yIKCLSWAHA", false),
				lesson(4, 4, "Overflow", "g0j0XaVk2EI", false),
				lesson(4, 5, "Tipo string", "AcyS0_BAy7U", false),
				lesson(4, 6, "Sistemas numericos", "Ma3M7Pdd7HI", false),
				lesson(4, 7, "Constantes", "Yaw80pKukMc", false),
				lesson(4, 8, "Iota", "1IduyaGMO3g", false),
				lesson(4, 9, "Deslocamento de bits", "USntMXkOihY", false),
			},
		},
		{
			Number: 5, Title: "Exercicios: Nivel #2", IsExercise: true, RequiredBy: 4,
			Lessons: []models.Lesson{
				lesson(5, 1, "Exercicio 1", "onK_nd--g1g", true),
				lesson(5, 2, "Exercicio 2", "8zEZw19gRTo", true),
				lesson(5, 3, "Exercicio 3", "12o81rTFu_w", true),
				lesson(5, 4, "Exercicio 4", "NxntmGW62Ag", true),
				lesson(5, 5, "Exercicio 5", "LSKCaxazLGs", true),
				lesson(5, 6, "Exercicio 6", "M07eDn7FyxI", true),
				lesson(5, 7, "Exercicio 7", "9eMbGsKcWlc", true),
			},
		},
		{
			Number: 6, Title: "Fluxo de Controle",
			Lessons: []models.Lesson{
				lesson(6, 1, "Entendendo fluxo de controle", "1G-tbQ6UE_A", false),
				lesson(6, 2, "Loops: inicializacao, condicao, pos", "g_Qdxi1b2cE", false),
				lesson(6, 3, "Loops: nested loop", "EL9wo6Zrz9o", false),
				lesson(6, 4, "Loops: a declaracao for", "QDaiqhTq3TA", false),
				lesson(6, 5, "Loops: break & continue", "eGdB8FMCZ0s", false),
				lesson(6, 6, "Loops: utilizando ascii", "hu0qcmbhH8s", false),
				lesson(6, 7, "Condicionais: a declaracao if", "LBlrrV0iRKg", false),
				lesson(6, 8, "Condicionais: if, else if, else", "ZfCgoVaxMGE", false),
				lesson(6, 9, "Condicionais: a declaracao switch", "WFFidtqPfhk", false),
				lesson(6, 10, "Condicionais: switch pt. 2 & documentacao", "v6HnDiPyynA", false),
				lesson(6, 11, "Operadores logicos condicionais", "Y5HamAGQzUg", false),
			},
		},
		{
			Number: 7, Title: "Exercicios: Nivel #3", IsExercise: true, RequiredBy: 6,
			Lessons: []models.Lesson{
				lesson(7, 1, "Exercicio 1", "xXUTHJzJNgM", true),
				lesson(7, 2, "Exercicio 2", "w_JW1bWT08s", true),
				lesson(7, 3, "Exercicio 3", "2y7qV_9SsiI", true),
				lesson(7, 4, "Exercicio 4", "bWvpNLBSEKk", true),
				lesson(7, 5, "Exercicio 5", "pIVF9vf2wAc", true),
				lesson(7, 6, "Exercicio 6", "VsDnqhpFVZI", true),
				lesson(7, 7, "Exercicio 7", "haYnFVssUr4", true),
				lesson(7, 8, "Exercicio 8", "mm3JM2ZG0us", true),
				lesson(7, 9, "Exercicio 9", "XxQQOHHQCwY", true),
				lesson(7, 10, "Exercicio 10", "3DP6lzRLkdw", true),
			},
		},
		{
			Number: 8, Title: "Agrupamentos de Dados",
			Lessons: []models.Lesson{
				lesson(8, 1, "Array", "i_3O4ooSVCM", false),
				lesson(8, 2, "Slice: literal composta", "MMzTlWZ9Gjw", false),
				lesson(8, 3, "Slice: for range", "l6O8HWaoy_w", false),
				lesson(8, 4, "Slice: fatiando ou deletando", "G0rxcnojV_U", false),
				lesson(8, 5, "Slice: anexando a uma slice", "MbvABKiAABA", false),
				lesson(8, 6, "Slice: make", "IMO5_ancK9w", false),
				lesson(8, 7, "Slice: slice multi-dimensional", "o3yoYGWqrDE", false),
				lesson(8, 8, "Slice: a surpresa do array subjacente", "dRNNC7VpztE", false),
				lesson(8, 9, "Maps: introducao", "clobcQqAgos", false),
				lesson(8, 10, "Maps: range & deletando", "7a6I-GnqtSE", false),
			},
		},
		{
			Number: 9, Title: "Exercicios: Nivel #4", IsExercise: true, RequiredBy: 8,
			Lessons: []models.Lesson{
				lesson(9, 1, "Exercicio 1", "cjUFrS3hqgU", true),
				lesson(9, 2, "Exercicio 2", "xpeExQ5C5S8", true),
				lesson(9, 3, "Exercicio 3", "Wzc5VOjh-XQ", true),
				lesson(9, 4, "Exercicio 4", "8Lym_4_dwOQ", true),
				lesson(9, 5, "Exercicio 5", "gA1cTnWjPaU", true),
				lesson(9, 6, "Exercicio 6", "KOhQCm8f8AE", true),
				lesson(9, 7, "Exercicio 7", "f4lvrtXsGIM", true),
				lesson(9, 8, "Exercicio 8", "rfQkR1bH3qw", true),
				lesson(9, 9, "Exercicio 9", "AKPdowl7tsw", true),
				lesson(9, 10, "Exercicio 10", "2cx79-nwNQU", true),
			},
		},
		{
			Number: 10, Title: "Structs",
			Lessons: []models.Lesson{
				lesson(10, 1, "Struct", "EaOGcmXo4F8", false),
				lesson(10, 2, "Structs embutidos", "KBFprVi_haM", false),
				lesson(10, 3, "Lendo a documentacao", "dKaElWlGKo0", false),
				lesson(10, 4, "Structs anonimos", "Y4MKS3gJQ9Q", false),
			},
		},
		{
			Number: 11, Title: "Exercicios: Nivel #5", IsExercise: true, RequiredBy: 10,
			Lessons: []models.Lesson{
				lesson(11, 1, "Exercicio 1", "OQMOIcZ-ShY", true),
				lesson(11, 2, "Exercicio 2", "uDajGJbXP6A", true),
				lesson(11, 3, "Exercicio 3", "ji14zPQgmN8", true),
				lesson(11, 4, "Exercicio 4", "uNnLSuRqBuY", true),
			},
		},
		{
			Number: 12, Title: "Funcoes",
			Lessons: []models.Lesson{
				lesson(12, 1, "Sintaxe", "PPOBe49M8V0", false),
				lesson(12, 2, "Desenrolando uma slice", "pV5OEqIRPh4", false),
				lesson(12, 3, "Defer", "zxVqQ59YfSY", false),
				lesson(12, 4, "Metodos", "YrCt7Wo2aAQ", false),
				lesson(12, 5, "Interfaces & polimorfismo", "2zTENBJTlD0", false),
				lesson(12, 6, "Funcoes anonimas", "pp8NKaoyefQ", false),
				lesson(12, 7, "Func como expressao", "j9C66R4BMWM", false),
				lesson(12, 8, "Retornando uma funcao", "9Oxmya_A-Sc", false),
				lesson(12, 9, "Callback", "u8qBzOAhbRM", false),
				lesson(12, 10, "Closure", "mOM0qTB5ppU", false),
				lesson(12, 11, "Recursividade", "1-pop5h5RAs", false),
			},
		},
		{
			Number: 13, Title: "Exercicios: Nivel #6", IsExercise: true, RequiredBy: 12,
			Lessons: []models.Lesson{
				lesson(13, 1, "Exercicio 1", "Uuj3-PrTk-Q", true),
				lesson(13, 2, "Exercicio 2", "ZLOyRrhSydk", true),
				lesson(13, 3, "Exercicio 3", "EoaMCgbqKSI", true),
				lesson(13, 4, "Exercicio 4", "HxdO5sm3cRY", true),
				lesson(13, 5, "Exercicio 5", "3pbmLfcgN2s", true),
				lesson(13, 6, "Exercicio 6", "NAXWCqhJIEU", true),
				lesson(13, 7, "Exercicio 7", "ePh12R5jnIM", true),
				lesson(13, 8, "Exercicio 8", "JS5G5O4LOVI", true),
				lesson(13, 9, "Exercicio 9", "EXQknIy6NE0", true),
				lesson(13, 10, "Exercicio 10", "LDJF2ceCRDE", true),
				lesson(13, 11, "Exercicio 11", "cOOH6mKWNs8", true),
			},
		},
		{
			Number: 14, Title: "Ponteiros",
			Lessons: []models.Lesson{
				lesson(14, 1, "O que sao ponteiros?", "l2YJ-5GpGr8", false),
				lesson(14, 2, "Quando usar ponteiros", "0slBes2RYgc", false),
			},
		},
		{
			Number: 15, Title: "Exercicios: Nivel #7", IsExercise: true, RequiredBy: 14,
			Lessons: []models.Lesson{
				lesson(15, 1, "Exercicio 1", "lSAVf0RgmHc", true),
				lesson(15, 2, "Exercicio 2", "XVd-y0t5fno", true),
			},
		},
		{
			Number: 16, Title: "Aplicacoes",
			Lessons: []models.Lesson{
				lesson(16, 1, "Documentacao JSON", "jnnIgvV0_yA", false),
				lesson(16, 2, "JSON marshal", "-tU2PSY8F5w", false),
				lesson(16, 3, "JSON unmarshal", "mcbj-wy8Ro8", false),
				lesson(16, 4, "A interface Writer", "S4hEdA0RPVI", false),
				lesson(16, 5, "O pacote sort", "b67JIGYM6Hc", false),
				lesson(16, 6, "Customizando o sort", "0E-q22d3QD4", false),
				lesson(16, 7, "bcrypt", "4vCb7jmwkzM", false),
			},
		},
		{
			Number: 17, Title: "Exercicios: Nivel #8", IsExercise: true, RequiredBy: 16,
			Lessons: []models.Lesson{
				lesson(17, 1, "Exercicio 1", "-wHOZpi697M", true),
				lesson(17, 2, "Exercicio 2", "oUsTxBwHaMM", true),
				lesson(17, 3, "Exercicio 3", "Y1Ym6Ai3uys", true),
				lesson(17, 4, "Exercicio 4", "fZZclCKnr7k", true),
				lesson(17, 5, "Exercicio 5", "3hidzcEZ0KE", true),
			},
		},
		{
			Number: 18, Title: "Concorrencia",
			Lessons: []models.Lesson{
				lesson(18, 1, "Concorrencia vs. paralelismo", "unofca9ooS4", false),
				lesson(18, 2, "Goroutines & WaitGroups", "4jXSU2jw3Ag", false),
				lesson(18, 3, "Discussao: Condicao de corrida", "0qGILXmLfMM", false),
				lesson(18, 4, "Na pratica: Condicao de corrida", "XxG7qqJzDKk", false),
				lesson(18, 5, "Mutex", "egd4WHJMwC0", false),
				lesson(18, 6, "Atomic", "iFlQ2yAYcp4", false),
			},
		},
		{
			Number: 19, Title: "Ambiente de Desenvolvimento",
			Lessons: []models.Lesson{
				lesson(19, 1, "O terminal", "CgrYtlaOgCg", false),
				lesson(19, 2, "Go Workspace", "geLSoE5D0xA", false),
				lesson(19, 3, "IDEs", "1E3KMnYnqCk", false),
				lesson(19, 4, "Comandos Go", "kJNq487dt7w", false),
				lesson(19, 5, "Repositorios no GitHub", "XQUMNuGyVa4", false),
				lesson(19, 6, "Explorando o GitHub", "t-S2s1K3b5Q", false),
				lesson(19, 7, "Compilacao cruzada", "k0pV_JoiZbI", false),
				lesson(19, 8, "Pacotes", "SO-RFPSqD3c", false),
			},
		},
		{
			Number: 20, Title: "Exercicios: Nivel #9", IsExercise: true, RequiredBy: 18,
			Lessons: []models.Lesson{
				lesson(20, 1, "Exercicio 1", "Y9QEvz4D_9E", true),
				lesson(20, 2, "Exercicio 2", "iu632z7i3MM", true),
				lesson(20, 3, "Exercicio 3", "cCWvFijhObU", true),
				lesson(20, 4, "Exercicio 4", "q_tHbwD0n6w", true),
				lesson(20, 5, "Exercicio 5", "58_JeZA3V_0", true),
				lesson(20, 6, "Exercicio 6", "weEJtEyl79o", true),
				lesson(20, 7, "Exercicio 7", "t0jraJONVzs", true),
			},
		},
		{
			Number: 21, Title: "Canais",
			Lessons: []models.Lesson{
				lesson(21, 1, "Entendendo canais", "jF0xuhnPkDg", false),
				lesson(21, 2, "Canais direcionais & Utilizando canais", "vYYHoKLb_8I", false),
				lesson(21, 3, "Range e close", "B1UArMoYDJ0", false),
				lesson(21, 4, "Select", "dp8s5jAc7h0", false),
				lesson(21, 5, "A expressao comma ok", "wWQ0BbbQ-28", false),
				lesson(21, 6, "Convergencia", "VJyryKEMleU", false),
				lesson(21, 7, "Divergencia", "8X6eOnSJu5g", false),
				lesson(21, 8, "Context", "PhTtrrsUH8c", false),
			},
		},
	}
}

// AllModules returns chapters grouped into learning modules.
func AllModules() []models.Module {
	chapters := AllChapters()
	chapterMap := make(map[int]models.Chapter)
	for _, ch := range chapters {
		chapterMap[ch.Number] = ch
	}

	return []models.Module{
		{
			Name: "Iniciante",
			Chapters: []models.Chapter{
				chapterMap[1], chapterMap[2], chapterMap[3],
				chapterMap[4], chapterMap[5], chapterMap[6], chapterMap[7],
			},
		},
		{
			Name: "Intermediario",
			Chapters: []models.Chapter{
				chapterMap[8], chapterMap[9], chapterMap[10], chapterMap[11],
				chapterMap[12], chapterMap[13], chapterMap[14], chapterMap[15],
				chapterMap[16], chapterMap[17],
			},
		},
		{
			Name: "Avancado",
			Chapters: []models.Chapter{
				chapterMap[18], chapterMap[20], chapterMap[21],
			},
		},
		{
			Name: "Ecossistema",
			Chapters: []models.Chapter{
				chapterMap[19],
			},
		},
	}
}
