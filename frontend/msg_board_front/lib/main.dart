import 'package:flutter/material.dart';
import 'dart:convert';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter App',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: const MyHomePage(),
    );
  }
}

class Message {
  final int id;
  final String name;
  final String text;
  final DateTime createdAt;

  Message (
    this.id,
    this.name,
    this.text,
    this.createdAt
  );

  Message.fromJson(Map<String, dynamic> json):
    id = json['id'] as int,
    name = json['name'] as String,
    text = json['text'] as String,
    createdAt = DateTime.parse(json['created_at'] as String);
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key});

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Home Page'),
      ),
      body: const Center(
        child: Text('Hello, World!'),
      ),
    );
  }
}