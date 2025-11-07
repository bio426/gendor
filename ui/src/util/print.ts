import { PDFDocument, StandardFonts, rgb } from "pdf-lib"

type OrderPayload = {
	id: number
	date: string
	plate: string
	propietary: string
	brand: string
	model: string
	year: string
	mileage: string
	observation: string
	discount: string
	items: {
		description: string
		price: string
	}[]
	subtotal: string
	total: string
}

function ajustarTexto(texto: string, limite: number) {
	if (texto.length <= limite) {
		return texto
	}

	let resultado = ""
	let posicionActual = 0

	while (posicionActual < texto.length) {
		let finLinea = posicionActual + limite
		if (finLinea >= texto.length) {
			resultado += texto.substring(posicionActual)
			break
		}

		let subCadena = texto.substring(posicionActual, finLinea)
		let ultimaPosicionEspacio = subCadena.lastIndexOf(" ")

		if (ultimaPosicionEspacio === -1) {
			// No hay espacio en el límite, hay que cortar la palabra
			resultado += texto.substring(posicionActual, finLinea) + "\n"
			posicionActual = finLinea
		} else {
			resultado +=
				texto.substring(
					posicionActual,
					posicionActual + ultimaPosicionEspacio
				) + "\n"
			posicionActual += ultimaPosicionEspacio + 1 // +1 para saltar el espacio
		}
	}
	return resultado
}

function ajustarTextoItems(texto: string, limite: number): [string, number] {
	if (texto.length <= limite) {
		return [texto, 1]
	}

	let resultado = ""
	let posicionActual = 0
	let linea = 1

	while (posicionActual < texto.length) {
		let finLinea = posicionActual + limite
		if (finLinea >= texto.length) {
			resultado += texto.substring(posicionActual)
			break
		}

		let subCadena = texto.substring(posicionActual, finLinea)
		let ultimaPosicionEspacio = subCadena.lastIndexOf(" ")

		if (ultimaPosicionEspacio === -1) {
			// No hay espacio en el límite, hay que cortar la palabra
			resultado += texto.substring(posicionActual, finLinea) + "\n"
			posicionActual = finLinea
		} else {
			resultado +=
				texto.substring(
					posicionActual,
					posicionActual + ultimaPosicionEspacio
				) + "\n"
			posicionActual += ultimaPosicionEspacio + 1 // +1 para saltar el espacio
		}
		linea += 1
	}
	return [resultado, linea]
}

export default {
	async printWithPdf(payload: OrderPayload) {
		const pdfRes = await fetch("/order.pdf")
		const pdfBuffer = await pdfRes.arrayBuffer()
		const pdfDoc = await PDFDocument.load(pdfBuffer)
		const helveticaFont = await pdfDoc.embedFont(StandardFonts.Helvetica)
		const page = pdfDoc.getPage(0)

		// fecha
		page.drawText(payload.date, {
			x: 510,
			y: 763,
			size: 12,
			font: helveticaFont,
		})
		// numeracion
		page.drawText(payload.id.toString().padStart(8, "0"), {
			x: 516,
			y: 738,
			size: 12,
			font: helveticaFont,
		})
		// propietario
		page.drawText(payload.propietary, {
			x: 110,
			y: 683,
			size: 12,
			font: helveticaFont,
		})
		// modelo
		page.drawText(payload.model, {
			x: 84,
			y: 663,
			size: 12,
			font: helveticaFont,
		})
		// placa
		page.drawText(payload.plate, {
			x: 74,
			y: 643,
			size: 12,
			font: helveticaFont,
		})
		// marca
		page.drawText(payload.brand, {
			x: 350,
			y: 683,
			size: 12,
			font: helveticaFont,
		})
		// anio
		page.drawText(payload.year.toString(), {
			x: 332,
			y: 663,
			size: 12,
			font: helveticaFont,
		})
		// kilometraje
		page.drawText(payload.mileage, {
			x: 386,
			y: 643,
			size: 12,
			font: helveticaFont,
		})
		// items
		let lastHeight = 570
		for (let i = 0; i < payload.items.length; i++) {
			const lineMargin = 10
			const clearText = payload.items[i].description.replace(/[\n\r\t]/gm, "")
			const [text, lineCount] = ajustarTextoItems(
				clearText,
				95
			)
			page.drawText(text, {
				x: 27,
				y: lastHeight,
				size: 10,
				font: helveticaFont,
				lineHeight: 12,
			})
			page.drawText(payload.items[i].price, {
				x: 500,
				y: lastHeight,
				size: 10,
				font: helveticaFont,
			})

			lastHeight = lastHeight - (lineCount * 12) - lineMargin
		}

		// observaciones
		page.drawText(ajustarTexto(payload.observation, 55), {
			x: 27,
			y: 137,
			size: 10,
			font: helveticaFont,
			lineHeight: 12,
		})

		// base text position for price box
		const textPosX = 568
		// subtotal
		const subtotalTextWidth = helveticaFont.widthOfTextAtSize(payload.subtotal, 12)
		page.drawText(payload.subtotal, {
			x: textPosX - subtotalTextWidth,
			y: 155,
			size: 12,
			font: helveticaFont,
		})
		// descuento
		const discountTextWidth = helveticaFont.widthOfTextAtSize(payload.discount, 12)
		page.drawText(payload.discount, {
			x: textPosX - discountTextWidth,
			y: 137,
			size: 12,
			font: helveticaFont,
		})
		// total
		const totalText = `S/. ${payload.total}`
		const totalTextWidth = helveticaFont.widthOfTextAtSize(totalText, 12)

		page.drawText(totalText, {
			x: textPosX - totalTextWidth,
			y: 72,
			size: 12,
			font: helveticaFont,
		})

		const newPdfBytes = await pdfDoc.save()
		const pdfUrl = URL.createObjectURL(
			new Blob([newPdfBytes as BlobPart], { type: "application/pdf" })
		)

		const win = window.open(pdfUrl, "_blank")
		if (win == null) return
		win.focus()
	},
}
