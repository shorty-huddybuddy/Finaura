// Remove trailing slashes to prevent double slashes in URLs
export const API_URL = (process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080').replace(/\/+$/, '')
export const ML_API_URL = (process.env.NEXT_PUBLIC_ML_API_URL || 'http://localhost:8000').replace(/\/+$/, '')
export const MARKET_DATA_URL = process.env.NEXT_PUBLIC_MARKET_DATA_URL || 'https://api.marketdata.app/v1'
export const MARKET_DATA_API_KEY = process.env.NEXT_PUBLIC_MARKET_DATA_API_KEY || '679542ee6b35c4.80378588'
