import ReactDOM from 'react-dom/client';
import './index.css';
import App from './componentes/App/App';
import reportWebVitals from './reportWebVitals';
import * as React from "react"
import {Provider} from "react-redux";
import store from "./store";
import {GoogleOAuthProvider} from "@react-oauth/google";

const root = ReactDOM.createRoot(
    document.getElementById('root') as HTMLElement
);
root.render(
    <Provider store={store}>
                <GoogleOAuthProvider clientId="40230221374-5jlod1u5abb0j1qjhd3frucpboadikda.apps.googleusercontent.com">
                    <App/>
                </GoogleOAuthProvider>
            </Provider>
);

reportWebVitals();
