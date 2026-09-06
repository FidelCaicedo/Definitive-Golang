---
description: Abre o cierra una sesión de estudio (repaso → objetivo → código → bitácora → commit).
argument-hint: abrir | cerrar
---

# Sesión de estudio

Modo: $ARGUMENTS — si viene vacío, deduce por el contexto: sin trabajo hecho en la
conversación, es `abrir`; con código escrito hoy, es `cerrar`. Ante duda, pregunta.

---

## Modo `abrir`

El enemigo es el día cero, no la sesión corta. Tu trabajo aquí es eliminar la fricción
de arranque, no añadir preámbulo. Máximo 15 minutos hasta que Fidel esté escribiendo Go.

1. **Repaso.** Ejecuta el protocolo de `.claude/commands/repaso.md`. Es innegociable y
   va primero: se recupera antes de leer nada nuevo.

2. **Estado real.** Lee la última entrada de `JOURNEY.md`, mira `git log --oneline -5` y
   el árbol de trabajo. Si la última entrada dejó una **pregunta abierta** en "Qué
   sigue", esa pregunta es la primera candidata al objetivo de hoy.

3. **Objetivo único.** Contrasta con la tabla de `PLAN.md` y propón **un solo**
   objetivo, formulado como algo que Fidel *escribe*, no que *lee*. Debe caber en
   25-30 minutos como piso. Si él propone otro, cede: la autonomía sostiene el hábito
   mejor que el plan.

4. **Fuente.** Nombra el capítulo concreto de `docs/books/` que sostiene el objetivo, con
   ruta de archivo. No resumas el capítulo por él.

5. **Contrato de la sesión.** Recuérdale que a partir de aquí rige la escalera de
   `/hint`: predicción primero, un peldaño a la vez. Y arranca.

Si Fidel vuelve tras una ausencia larga, no la conviertas en tema. Una frase de encuadre
y al código: el hábito se restaura ejecutándolo, no procesándolo.

---

## Modo `cerrar`

1. **Verificación técnica.** Antes de celebrar nada, corre sobre el módulo tocado:

   ```
   gofmt -l . && go vet ./... && go build ./...
   ```

   Reporta el resultado tal cual. Si falla algo, se arregla antes de cerrar.

2. **Revisión con estándar profesional.** Idiomatismo, manejo de errores, nombres,
   qué se copia y qué no. Señala lo que un revisor de Go marcaría — y **nombra
   explícitamente lo que hizo bien**, sobre todo si generalizó más allá de lo pedido.

3. **Teach-back.** Antes de escribir la bitácora, pídele que explique en sus palabras
   la mecánica central de hoy. Lo que no pueda explicar no lo aprendió: eso va a
   "qué costó", no a "qué aprendí".

4. **Bitácora.** Escribe la entrada nueva **al principio** de `JOURNEY.md` (orden
   cronológico inverso), con fecha real y la estructura existente:
   - **Qué se hizo** — el artefacto concreto, con módulo y estado de las verificaciones.
   - **Qué se aprendió** — mecánicas específicas, no temas. "`strings.Fields` colapsa
     espacios vía `unicode.IsSpace`", no "vi strings".
   - **Qué costó** — incluye las autopsias de `/autopsy` y lo que falló en el repaso.
   - **Qué sigue** — termina con **una pregunta abierta y concreta** para la próxima
     sesión. Es lo que hace posible el `/repaso` de la próxima vez.

   Preserva las entradas históricas intactas. No reescribas errores pasados.

5. **Commit.** Propón el mensaje siguiendo el estilo del historial (`git log`), muéstralo
   y espera aprobación antes de commitear. **Pregunta antes de hacer push** o de abrir
   un PR — nunca lo asumas.

6. **Cierre.** Una línea sobre el terreno ganado. Sin inflarlo: la maestría se acumula
   en sesiones que no se sintieron memorables.
