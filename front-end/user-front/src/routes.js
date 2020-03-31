import React from 'react';
import  {BrowserRouter, Switch, Route} from 'react-router-dom';

import CadastroPessoa from './pages/cadastroPessoa';
import ListaPessoa from './pages/listaPessoa';


export default function Routes(){
    return(
        <BrowserRouter>
        <Switch>
            <Route path="/" exact component={CadastroPessoa}/>
            <Route path="/listar" component={ListaPessoa}/>
        </Switch>
        </BrowserRouter>
    )
}