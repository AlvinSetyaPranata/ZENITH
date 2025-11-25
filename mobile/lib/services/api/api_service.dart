import 'dart:convert';
import 'package:http/http.dart' as http;

class ApiService {
  static const String baseUrl = 'http://localhost:8000/api';

  // Example method for making API calls
  Future<http.Response> getData(String endpoint) async {
    final response = await http.get(Uri.parse('$baseUrl/$endpoint'));
    return response;
  }

  Future<http.Response> postData(String endpoint, dynamic data) async {
    final response = await http.post(
      Uri.parse('$baseUrl/$endpoint'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(data),
    );
    return response;
  }

  Future<http.Response> putData(String endpoint, dynamic data) async {
    final response = await http.put(
      Uri.parse('$baseUrl/$endpoint'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(data),
    );
    return response;
  }

  Future<http.Response> deleteData(String endpoint) async {
    final response = await http.delete(Uri.parse('$baseUrl/$endpoint'));
    return response;
  }
}