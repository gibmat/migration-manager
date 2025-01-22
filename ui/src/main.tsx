import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter } from 'react-router';
import {
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query';
import './css/index.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import App from './App.tsx';
import { AuthProvider } from 'context/auth';
import { NotificationProvider } from 'context/notification';

const queryClient = new QueryClient();

createRoot(document.getElementById('root')!).render(
  <QueryClientProvider client={queryClient}>
    <AuthProvider>
      <BrowserRouter>
        <NotificationProvider>
          <StrictMode>
            <App />
          </StrictMode>
        </NotificationProvider>
      </BrowserRouter>
    </AuthProvider>
  </QueryClientProvider>
);
