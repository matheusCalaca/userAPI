import React from 'react'
import { Component } from "react";
import { Scope } from '@unform/core'
import Input from '../Input';

class EnderecoForm extends Component {

    state = {
        enderecoContador: 1,
    }


    handlePlusEndereco = (event) => {
        event.preventDefault(); 
        this.setState({
            enderecoContador: this.state.enderecoContador + 1
        })
    }

    render() {
        let formTelefone = [];
        for (var i = 0; i < this.state.enderecoContador; i++) {
            formTelefone.push(
                <div key={i}>
                    <h1>Enderco {i}</h1>
                    <Scope path={"enderecos[" + i + "]"} >
                        <span>CEP: </span> <br />
                        <Input name="CEP" type="number" /> <br />
                        <span>logradouro: </span> <br />
                        <Input name="logradouro" type="text" /> <br />
                        <span>bairro: </span> <br />
                        <Input name="bairro" type="text" /> <br />
                        <span>cidade: </span> <br />
                        <Input name="cidade" type="text" /> <br />
                        <span>UF: </span> <br />
                        <Input name="UF" type="text" /> <br />
                        <span>complemento: </span> <br />
                        <Input name="complemento" type="text" /> <br />
                        <span>numero: </span> <br />
                        <Input name="numero" type="text" /> <br />
                        <span>tipo: </span> <br />
                        <Input name="tipo" type="text" /> <br />
                    </Scope>
                </div>
            );
        }

        return (
            <>
                {formTelefone}
                <button onClick={this.handlePlusEndereco}>End ++</button><br /><br />

            </>
        );
    }
}
export default EnderecoForm;