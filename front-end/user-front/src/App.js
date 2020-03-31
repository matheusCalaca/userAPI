import React, { useState } from 'react';
import './App.css';
// import api from './services/app'

function App() {

  const [nome, setNome] = useState('');

  function handleSubmit(event) {
    event.preventDefault()
    console.log('Nome: ' + nome);

    // values


  }



  return (
    <div className="container">
      <div className="container">

        <h1>Cadastreo de Pessoa</h1>
        <form onSubmit={handleSubmit}>
          <label htmlFor="nome">Nome:</label>
          <input
            type="text"
            id="nome"
            placeholder="Digite seu Nome"
            value={nome}
            onChange={event => setNome(event.target.value)}
          />
          <button type="submit">Entrar</button>
        </form>
      </div>
    </div>
  );
}

export default App;
