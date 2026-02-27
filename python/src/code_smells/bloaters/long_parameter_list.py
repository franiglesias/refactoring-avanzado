from datetime import datetime


class ReportGenerator:
    def generate_report(
        self,
        title: str,
        start_date: datetime,
        end_date: datetime,
        include_charts: bool,
        include_summary: bool,
        author_name: str,
        author_email: str,
    ):
        print(f"Generando reporte: {title}")
        print(f"Desde {start_date.strftime('%x')} hasta {end_date.strftime('%x')}")
        print(f"Autor: {author_name} ({author_email})")
        if include_charts:
            print('Incluyendo gráficos...')
        if include_summary:
            print('Incluyendo resumen...')
        print('Reporte generado exitosamente.')


def demo_long_parameter_list():
    gen = ReportGenerator()
    gen.generate_report(
        'Ventas Q1',
        datetime(2025, 1, 1),
        datetime(2025, 3, 31),
        True,
        False,
        'Pat Smith',
        'pat@example.com',
    )
