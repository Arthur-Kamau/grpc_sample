import { Component, createEffect } from 'solid-js';
import { getTokenStorage } from './util/localStorage';
import {appState, setAppState} from "./store/AppStore";
import {Route} from "@solidjs/router";
import {GrpcWebFetchTransport} from "@protobuf-ts/grpcweb-transport";
import {AppApiClient} from "./gen/app.client";
import  toast  from 'solid-toast';
import DashboardPage from './modules/dashboard/dashboard';
import HomePage from './modules/home/home';

const AppRoutes: Component = () => {
   const API_URL = 'http://127.0.0.1:8080';

  createEffect(async () => {
    try {
        const token = getTokenStorage()
        setAppState('token',token??"")


        if(appState.client==null){
            const transport = new GrpcWebFetchTransport({
                baseUrl: API_URL,
            });
            const client = new AppApiClient(transport);

            setAppState('client', client)


        }

      
// todo navigate to ligin if token is invalid


    } catch (error) {
        toast.error("Error while loading grpc client");
        console.log(error);
    }
})
  return (

      <>
           <Route path="/" component={DashboardPage}/>
           <Route path="/home" component={HomePage}/>
           </>

  );
};

export default AppRoutes;
