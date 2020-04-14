import React from 'react';
import  {BrowserRouter, Switch, Route} from 'react-router-dom';

import CadastroPessoa from './pages/cadastroPessoa';
import listarPessoa from './pages/pessoa/listarPessoa';


export default function Routes(){
    return(
        <BrowserRouter>
        <Switch>
            <Route path="/" exact component={listarPessoa}/>
            <Route path="/cadastrar" component={CadastroPessoa}/>
            <Route path="/listar" component={listarPessoa}/>
        </Switch>
        </BrowserRouter>
    )
}