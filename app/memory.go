package main

import "sync"

var (
	mutexStudent sync.Mutex
	studentToMarks map[string][]float64	
)

var (
	mutexExam sync.Mutex
	examToMarks map[string][]float64	
)
