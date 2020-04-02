import React from 'react'
import { Component } from "react";
import { Scope } from '@unform/core'
import Input from '../Input';

class PhoneForm extends Component {

    state = {
        phoneContador: 1,
    }

    handlePlusTelefone = (event) => {
        event.preventDefault(); 
        this.setState({
            phoneContador: this.state.phoneContador + 1,
        })
    }

    render() {
        let formTelefone = [];
        for (var i = 0; i < this.state.phoneContador; i++) {
            formTelefone.push(
                <div key={i}>
                    <h1>Telefone {i}</h1>
                    <Scope path={"telefones[" + i + "]"} >
                        <span>ddd: </span> <br />
                        <Input name="ddd" type="number" /> <br />
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
                <button onClick={this.handlePlusTelefone}>Tel ++</button><br /><br />
            </>
        );
    }
}
export default PhoneForm;