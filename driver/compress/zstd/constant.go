package zstd

// 編碼併發處理限制 不可小於 1
// 1: 單執行續
// GOMAXPROCS: 建議數量 *預設
// const WithEncoderConcurrencySize =

// 解碼併發處理限制
// 0: 限制等同 cpu 數量 *預設
// 1: 單執行續
// 4: 建議最大數量
// const WithDecoderConcurrencySize =

// 編碼單次處理最大容量
// 預設: 1<<29 : 512M
const WithWindowSize = 1 << 20 // 1M

// 解碼單次處理最大容量
// 預設: 1<<29 : 512M
const WithDecoderMaxWindowSize = 1 << 20 // 1M

// 解碼整理使用最大容量
// 預設: 64 << 30 : 64G
const WithDecoderMaxMemorySize = 1024 << 20 // 1G
