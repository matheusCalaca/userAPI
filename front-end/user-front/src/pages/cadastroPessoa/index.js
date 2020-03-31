import React, { useState }  from 'react';


export default function CadastroPessoa() {
    const [nome, setNome] = useState('');

    function handleSubmit(event) {
      event.preventDefault()
      console.log('Nome: ' + nome);
  
      // values
  
  
    }
  
  
  

    return (
        <>
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
        </>
    )
}