import ReactDOM from 'react-dom/client';
import './index.css';
import App from './componentes/App/App';
import reportWebVitals from './reportWebVitals';
import * as React from "react"
import {Provider} from "react-redux";
import store from "./store";

const root = ReactDOM.createRoot(
    document.getElementById('root') as HTMLElement
);
root.render(
    <React.StrictMode>
        <Provider store={store}>
            <App/>
        </Provider>
    </React.StrictMode>
);

reportWebVitals();
