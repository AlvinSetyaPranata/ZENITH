import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:mobile/services/api/api_service.dart';

class AuthService {
  final ApiService _apiService = ApiService();

  Future<bool> login(String email, String password) async {
    try {
      final response = await _apiService.postData('auth/login', {
        'email': email,
        'password': password,
      });

      if (response.statusCode == 200) {
        // Handle successful login
        return true;
      } else {
        // Handle login error
        return false;
      }
    } catch (e) {
      // Handle network error
      return false;
    }
  }

  Future<bool> register(Map<String, dynamic> userData) async {
    try {
      final response = await _apiService.postData('auth/register', userData);

      if (response.statusCode == 201) {
        // Handle successful registration
        return true;
      } else {
        // Handle registration error
        return false;
      }
    } catch (e) {
      // Handle network error
      return false;
    }
  }

  Future<void> logout() async {
    // Implementation for logout
    // This would typically clear user session/token
  }
}