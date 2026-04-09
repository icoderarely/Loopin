import { useState } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import './ConfirmationPage.css'

const API_URL = 'http://localhost:8080/v1'

export function ConfirmationPage() {
  const { token = '' } = useParams()
  const navigate = useNavigate()
  const [status, setStatus] = useState<'idle' | 'loading' | 'success' | 'error'>('idle')
  const [errorMessage, setErrorMessage] = useState('')

  const handleConfirm = async () => {
    setStatus('loading')
    setErrorMessage('')

    try {
      const response = await fetch(`${API_URL}/users/activate/${token}`, {
        method: 'PUT',
      })

      if (response.ok) {
        setStatus('success')
        setTimeout(() => navigate('/'), 3000)
      } else {
        setStatus('error')
        setErrorMessage('Failed to activate account. The link may have expired.')
      }
    } catch {
      setStatus('error')
      setErrorMessage('Network error. Please try again.')
    }
  }

  return (
    <div className="confirmation-container">
      <div className="confirmation-card">
        {status === 'idle' && (
          <>
            <div className="icon-container">
              <svg className="mail-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
              </svg>
            </div>
            <h1 className="title">Confirm Your Email</h1>
            <p className="description">
              Click the button below to activate your Loopin account and get started.
            </p>
            <button className="confirm-button" onClick={handleConfirm}>
              Activate Account
            </button>
          </>
        )}

        {status === 'loading' && (
          <>
            <div className="spinner"></div>
            <h1 className="title">Activating...</h1>
            <p className="description">Please wait while we verify your account.</p>
          </>
        )}

        {status === 'success' && (
          <>
            <div className="icon-container success">
              <svg className="check-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 13l4 4L19 7" />
              </svg>
            </div>
            <h1 className="title success">Account Activated!</h1>
            <p className="description">Your account has been successfully activated.</p>
            <p className="redirect-text">Redirecting to home page...</p>
          </>
        )}

        {status === 'error' && (
          <>
            <div className="icon-container error">
              <svg className="error-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
              </svg>
            </div>
            <h1 className="title error">Activation Failed</h1>
            <p className="description">{errorMessage}</p>
            <button className="confirm-button retry" onClick={() => setStatus('idle')}>
              Try Again
            </button>
          </>
        )}
      </div>
    </div>
  )
}