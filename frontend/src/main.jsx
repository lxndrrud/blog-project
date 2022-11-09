import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import { Provider } from 'react-redux'
import { setupStore } from './store/store'

ReactDOM.createRoot(document.getElementById('root')).render(
    <Provider store={setupStore()} >
        <App />
    </Provider>
)
