import 'package:flutter/material.dart';
import 'package:mobile/components/shared/app_bar.dart';
import 'package:mobile/components/widgets/custom_button.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: const CustomAppBar(
        title: 'Home',
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Text(
              'Welcome to ZENITH Student Management System',
              textAlign: TextAlign.center,
              style: TextStyle(
                fontSize: 24,
                fontWeight: FontWeight.bold,
              ),
            ),
            const SizedBox(height: 20),
            CustomButton(
              text: 'Get Started',
              onPressed: () {
                // Handle button press
              },
            ),
          ],
        ),
      ),
    );
  }
}