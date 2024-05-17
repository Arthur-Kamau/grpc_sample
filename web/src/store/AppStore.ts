
import {createStore} from "solid-js/store";
import {AppApiClient} from '../gen/app.client';

const [appState, setAppState] = createStore({
    client: undefined as AppApiClient | undefined,
    token: '',
});


export {appState, setAppState}