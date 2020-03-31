import React from 'react';
import './App.css';
import api from './services/app'

function App() {
  return (
    <div className="container">
      <div className="container">

        <h1>Cadastreo de Pessoa</h1>
        <form htmlFor="cadastroPessoa">
          <label htmlFor="nome">Nome:</label>
          <input
            type="text"
            id="nome"
            placeholder="Digite seu Nome"
          />
          <button type="submit">Entrar</button>
        </form>
      </div>
    </div>
  );
}

export default App;
