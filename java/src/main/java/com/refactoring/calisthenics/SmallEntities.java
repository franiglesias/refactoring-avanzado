package com.refactoring.calisthenics;

import java.util.*;

public class SmallEntities {

    public static void main(String[] args) {
        ReportService service = new ReportService();
        String json = "[{\"name\":\"Ana\",\"age\":30},{\"name\":\"Luis\",\"age\":25}]";
        String csv = service.generateCsvReportFromJson(json, ",");
        System.out.println(csv);
    }

    public static class ReportService {
        private Map<String, String> cache = new HashMap<>();

        public String generateCsvReportFromJson(String jsonInput, String delimiter) {
            // Este método es demasiado largo - debería dividirse
            Object data;
            try {
                // Simulación de parsing JSON (en Java usarías una librería como Jackson)
                data = parseJsonArray(jsonInput);
            } catch (Exception e) {
                throw new IllegalArgumentException("Invalid JSON");
            }

            if (!(data instanceof List)) {
                throw new IllegalArgumentException("Expected array");
            }

            @SuppressWarnings("unchecked")
            List<Map<String, Object>> list = (List<Map<String, Object>>) data;

            if (list.isEmpty()) {
                return "";
            }

            List<String> headers = new ArrayList<>(list.get(0).keySet());
            List<String> lines = new ArrayList<>();
            lines.add(String.join(delimiter, headers));

            for (Map<String, Object> row : list) {
                List<String> values = new ArrayList<>();
                for (String h : headers) {
                    Object val = row.get(h);
                    values.add(val != null ? String.valueOf(val) : "");
                }
                lines.add(String.join(delimiter, values));
            }

            String result = String.join("\n", lines);
            cache.put("last", result);
            return result;
        }

        // Método auxiliar para simular parsing JSON
        private Object parseJsonArray(String json) {
            // Simulación simple - en producción usar Jackson o Gson
            List<Map<String, Object>> result = new ArrayList<>();
            if (json.contains("Ana")) {
                Map<String, Object> row1 = new HashMap<>();
                row1.put("name", "Ana");
                row1.put("age", 30);
                result.add(row1);

                Map<String, Object> row2 = new HashMap<>();
                row2.put("name", "Luis");
                row2.put("age", 25);
                result.add(row2);
            }
            return result;
        }
    }
}
