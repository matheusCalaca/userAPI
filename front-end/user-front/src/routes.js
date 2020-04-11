import React from 'react';
import  {BrowserRouter, Switch, Route} from 'react-router-dom';

import CadastroPessoa from './pages/cadastroPessoa';
import ListaPessoa from './pages/listaPessoa';
import Dashboard from './pages/dashboard/Dashboard';


export default function Routes(){
    return(
        <BrowserRouter>
        <Switch>
            <Route path="/" exact component={CadastroPessoa}/>
            <Route path="/teste" exact component={Dashboard}/>
            <Route path="/listar" component={ListaPessoa}/>
        </Switch>
        </BrowserRouter>
    )
}