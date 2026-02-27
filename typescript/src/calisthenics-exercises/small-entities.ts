export class ReportService {
  private cache: Map<string, string> = new Map()

  generateCsvReportFromJson(jsonInput: string, delimiter: string = ','): string {
    let data: unknown
    try {
      data = JSON.parse(jsonInput)
    } catch (e) {
      throw new Error('Invalid JSON')
    }

    if (!Array.isArray(data)) {
      throw new Error('Expected array')
    }

    const headers = Object.keys(data[0] as any)
    const lines = [headers.join(delimiter)]
    for (const row of data as Array<Record<string, any>>) {
      const values = headers.map((h) => String(row[h] ?? ''))
      lines.push(values.join(delimiter))
    }
    const result = lines.join('\n')

    this.cache.set('last', result)
    return result
  }
}
