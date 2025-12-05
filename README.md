# 📊 Finaura – GenAI-Powered Financial Assistant

Finaura is an AI-powered financial assistant that simplifies investment planning, portfolio management, and financial education. It empowers users with real-time insights, personalized suggestions, and smart financial tools — all in one seamless platform.

## 🌐 Live Demo

**Frontend**: [https://finaura-kohl.vercel.app/](https://finaura-kohl.vercel.app/)  
**Backend API**: [https://finaura-backend-xfvz.onrender.com](https://finaura-backend-xfvz.onrender.com)

---

## 🚩 Problem Statement

In today’s dynamic financial landscape, individuals often face:
- Information overload
- Lack of trustworthy investment advice
- Low financial literacy
- Difficulty managing personal finances

**Finaura** aims to bridge this gap with automation, education, and GenAI.

---

## 🌟 Key Features

### 🧠 AI-Powered Insights
- **Portfolio Evaluation** – Get BUY/HOLD/SELL actions
- **Risk Profiling** – Classify users as Low, Medium, or High risk
- **Smart Recommendations** – Tailored investment options
- **Market Trends** – AI-curated financial news and analysis

### 💬 TradeMate Chatbot
- Get instant answers to financial questions
- Investment suggestions and general help via GenAI

### 🎓 Financial Education Hub
- Blogs, videos, and micro-courses for every level

### 📅 Financial Calendar
- Track EMIs, tax deadlines, and investment reminders

### 📈 Market Monitoring
- Real-time asset tracking, personalized watchlists

### 🤝 Social Trading
- Follow top traders
- Learn strategies from experienced investors

### 💸 Pay-to-Earn Model
- Earn rewards for using features, completing tasks, and referring users

---

## 🧰 Tech Stack

| Layer        | Technology      |
|--------------|-----------------|
| Frontend     | Next.js         |
| Backend      | Go (Golang) /Python    |
| Database     | Firebase        |
| AI/LLM       | Gemini API/Models |

---

## 🚀 Business Model

- **Freemium**: Core features available for free
- **Premium**: Advanced analytics, premium social trading
- **Pay-to-Earn**: Incentives for education, referrals, and engagement

---

## 📽️ Demo Video

Watch the 3-minute product demo: [https://www.youtube.com/watch?v=E2v8UEyJjDU](https://www.youtube.com/watch?v=E2v8UEyJjDU)

---

## Project Structure
```
/finapp/
├── frontend/                   # Next.js frontend application
│   ├── src/
│   │   ├── app/                # Next.js app router routes
│   │   │   └── stock-dashboard/# Stock dashboard page
│   │   ├── components/         # Reusable React components
│   │   │   ├── ui/             # UI components (cards, tabs, etc.)
│   │   │   ├── *-widget.tsx    # TradingView widget integrations
│   │   │   └── ...
│   │   ├── constants/          # Application constants
│   │   │   └── education.ts    # Educational content data
│   │   ├── hooks/              # Custom React hooks
│   │   │   └── store.ts        # Transaction state management
│   │   ├── lib/                # Utility functions
│   │   └── types/              # TypeScript type definitions
│   ├── public/                 # Static assets
│   ├── cloudbuild.yaml         # Google Cloud Build configuration
│   └── next.config.ts          # Next.js configuration
│
└── backend/                    # Go backend application
    ├── config/                 # Configuration setup
    │   └── ...                 # Stripe and other service configs
    ├── database/               # Database integration
    │   └── ...                 # Firebase initialization
    ├── routes/                 # API route handlers
    │   └── ...
    └── main.go                 # Main application entry point
```



## Environment Variables

The application uses various environment variables for API keys and service configuration:

- `NEXT_PUBLIC_API_URL`: Backend API URL
- `NEXT_PUBLIC_ML_API_URL`: Machine learning recommendation service URL
- `NEXT_PUBLIC_MARKET_DATA_URL`: Market data API URL
- `NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY`: Stripe public key
- `CLERK_SECRET_KEY`: Clerk authentication key




## 📌 Getting Started (for Developers)

### Prerequisites
- Node.js 18+ and npm
- Go 1.23+
- Python 3.11+ (for ML recommendation service)
- Firebase account
- Clerk account
- Stripe account
- Google Gemini API key

### 1. Clone the Repository
```bash
git clone https://github.com/shorty-huddybuddy/Finaura.git
cd Finaura
```

### 2. Frontend Setup

#### Install Dependencies
```bash
cd frontend
npm install
```

#### Configure Environment Variables
Create a `.env.local` file in the `frontend/` directory:

```env
# Backend API
NEXT_PUBLIC_API_URL=http://localhost:8080

# ML Recommendation Service
NEXT_PUBLIC_ML_API_URL=http://localhost:8000

# Market Data
NEXT_PUBLIC_MARKET_DATA_URL=https://api.marketdata.app/v1
NEXT_PUBLIC_MARKET_DATA_API_KEY=your_market_data_api_key

# Clerk Authentication
NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY=your_clerk_publishable_key
CLERK_SECRET_KEY=your_clerk_secret_key

# Stripe Payment
NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY=your_stripe_publishable_key
```

#### Start Frontend Development Server
```bash
npm run dev
```
Frontend will run on: http://localhost:3000

### 3. Backend Setup (Go)

#### Configure Environment Variables
Create a `.env` file in the `backend/` directory:

```env
# Server Configuration
API_PORT=8080
APP_ENV=development
FRONTEND_URL=http://localhost:3000

# Clerk Authentication
CLERK_SECRET_KEY=your_clerk_secret_key

# Stripe Payment
STRIPE_SECRET_KEY=your_stripe_secret_key
STRIPE_WEBHOOK_SECRET=your_webhook_secret

# Google Gemini AI
GEMINI_API_KEY=your_gemini_api_key

# Finnhub Stock Market API
FINNHUB_API_KEY=your_finnhub_api_key

# Firebase
FIREBASE_DB_URL=https://your-project-default-rtdb.firebaseio.com
```

#### Add Firebase Credentials
Place your Firebase service account JSON file as `api_key.json` in the `backend/` directory.

#### Install Dependencies and Run
```bash
cd backend
go mod tidy
go run main.go
```
Backend will run on: http://localhost:8080

### 4. ML Recommendation Service (Python)

#### Install Dependencies
```bash
cd recommendations_model
pip install -r requirements.txt
```

#### Start Python Server
```bash
uvicorn app:app --reload
```
ML service will run on: http://localhost:8000

---

## 🚢 Deployment

- **Frontend**: Deployed on [Vercel](https://vercel.com) → [https://finaura-kohl.vercel.app/](https://finaura-kohl.vercel.app/)
- **Backend**: Deployed on [Render](https://render.com) → [https://finaura-backend-xfvz.onrender.com](https://finaura-backend-xfvz.onrender.com)

For detailed deployment instructions, see [SETUP_GUIDE.md](./SETUP_GUIDE.md)

---

## 📸 UI Screenshots

![image](https://github.com/user-attachments/assets/077bc3b0-5ed3-4098-941d-7caefcbb78c9)
![image](https://github.com/user-attachments/assets/cb36c013-13df-457e-ae5c-ed703754c4c8)
![image](https://github.com/user-attachments/assets/1d8194ab-d1b4-47df-8fec-672da6a77159)
![image](https://github.com/user-attachments/assets/89796cc6-d2ab-47de-9002-c487835f84df)
![image](https://github.com/user-attachments/assets/0511be0e-8f8f-4c7c-88d0-71404e4f29de)
![image](https://github.com/user-attachments/assets/262c94f4-3749-487c-b51c-01a670d71f81)
![image](https://github.com/user-attachments/assets/91b19b71-de09-42bc-b063-bd83a867d151)

