# Plan de estudio — de cero a artesano de Go

**Presupuesto real:** 3-5 h/semana. **Enfoque:** proyecto desde el día 1.
**Regla de oro:** cada sesión termina con código escrito por Fidel y commiteado, aunque sean 30 líneas.

## Proyecto norte

Motor de búsqueda full-text en Go (inverted index → BM25 → top-k), construido desde cero, pieza por pieza, conforme avanzan los capítulos. Elegido sobre el bandit engine por estar claramente fuera del dominio de MercadoLibre (sin riesgo de IP). El bandit engine puede ser el proyecto #2 si el contrato lo permite (revisar cláusulas de IP/moonlighting antes).

## Fase 1 — Fundamentos aplicados (~3 meses, Learning Go caps. 1-10)

Ritmo: ~1 capítulo por semana. Cada capítulo se aplica al proyecto:

| Semana | Learning Go | Aplicación al proyecto |
|---|---|---|
| 1 | Cap. 1-2: entorno, tipos | Módulo `01-fundamentos`; CLI mínima que lee un archivo de texto |
| 2 | Cap. 3: strings, slices, maps | Tokenizador: partir texto en palabras, contar frecuencias |
| 3 | Cap. 4: bloques, control | Normalización (minúsculas, puntuación); tabla de stopwords |
| 4 | Cap. 5: funciones | Pipeline de indexado como funciones componibles |
| 5 | Cap. 6: punteros | Entender qué se copia y qué no en el índice; benchmarks simples |
| 6 | Cap. 7: tipos, métodos, interfaces | `Index` como tipo con métodos; interfaz `Tokenizer` |
| 7 | Cap. 8: generics | Estructuras genéricas auxiliares (set, heap para top-k) |
| 8 | Cap. 9: errores | Manejo de errores del pipeline de indexado (archivos corruptos, etc.) |
| 9 | Cap. 10: módulos, paquetes | Reorganizar el proyecto en paquetes con API limpia |
| 10-12 | Repaso + 100 Go Mistakes (caps. relevantes) | Auditar el propio código contra los errores del libro |

## Fase 2 — Profundidad (~3 meses)

- Learning Go caps. 11-16 (stdlib, contexto, testing, reflection).
- Testing serio del motor: tabla-driven tests, benchmarks, `pprof`.
- Ranking BM25; consultas booleanas; persistencia del índice en disco.
- *100 Go Mistakes* como revisión continua.

## Fase 3 — Concurrencia (~2-3 meses)

- *Concurrency in Go* completo + Learning Go cap. 12.
- Indexado concurrente (fan-out de goroutines, channels, backpressure).
- Servidor HTTP de búsqueda.

## CLRS (Introduction to Algorithms)

Solo referencia *just-in-time*: se abre un capítulo cuando el proyecto lo exige (heaps para top-k, hashing, árboles). No es un track paralelo.

## Rituales

- Cierre de sesión: entrada en `JOURNEY.md` (qué aprendí, qué costó, qué sigue) + commit + push.
- Check-in semanal automatizado con Claude: revisión de progreso contra este plan.
- Sesión mínima viable: 25-30 min. El enemigo es el día cero, no la sesión corta.
