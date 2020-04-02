import React from 'react'
import { Component } from "react";
import { Socope, Scope } from '@unform/core'
import Input from '../Input';

class PhoneForm extends Component {



    render() {
        let formTelefone = [];
        for (var i = 0; i < this.props.teste; i++) {
            formTelefone.push(
                <>
                    <h1>Telefone {i}</h1>
                    <Scope path={"telefones[" + i + "]"} >
                        <span>ddd: </span> <br />
                        <Input name="ddd" type="number" /> <br />
                        <span>numero: </span> <br />
                        <Input name="numero" type="text" /> <br />
                        <span>tipo: </span> <br />
                        <Input name="tipo" type="text" /> <br />
                    </Scope>
                </>
            );
        }

        return (
            <>
                {formTelefone}
            </>
        );
    }
}
export default PhoneForm;