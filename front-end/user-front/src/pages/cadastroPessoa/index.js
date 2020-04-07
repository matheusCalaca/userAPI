import { Form } from '@unform/web';
import React, { Component } from 'react';
import EnderecoForm from '../../components/form/endereco-form';
import Input from '../../components/form/Input';
import PhoneForm from '../../components/form/phone-form';
import Axios from 'axios';
import { PessoaDto } from '../../models/pessoaDTO'




class CadastroPessoa extends Component {
    handleSubmit = (data) => {
        var pessoa = new PessoaDto();
        if (data != undefined) {
            pessoa = JSON.parse(JSON.stringify(data));
            pessoa.telefones.forEach(tel => { tel.ddd = parseInt(tel.ddd); });
            pessoa.dataNascimento = new Date(pessoa.dataNascimento);
        }
        const json = JSON.stringify(pessoa);
        Axios.post(`http://localhost:8000/api/users`, json, {
            headers: { "Content-Type": "application/json" }
        }).then(
            (resp) => {
                console.log(resp.data)
                alert('sucesso');
            }

        )
            .catch(
                (err) => {
                    console.log("ERRO - " + JSON.stringify(err.response.data));
                    console.log(err.response['data']);
                    alert("ERRO - " + JSON.stringify(err.response.data));
                }
            )
        console.log(JSON.stringify(pessoa));
    }

    render() {
        return (
            <Form onSubmit={this.handleSubmit}>
                <div className="contentForm">
                    <div className="contentFornInsider">
                        <div>
                            <span>nome: </span> <br />
                            <Input name="nome" type="text" /> <br />
                        </div>
                        <div>
                            <span>sebrenome: </span> <br />
                            <Input name="sobrenome" type="text" /> <br />
                        </div>
                        <div>
                            <span>CPF: </span> <br />
                            <Input name="CPF" type="text" /> <br />
                        </div>
                        <div>
                            <span>RG: </span> <br />
                            <Input name="RG" type="text" /> <br />
                        </div>
                        <div>
                            <span>dataNascimento: </span> <br />
                            <Input name="dataNascimento" type="date" /> <br />
                        </div>
                        <div>
                            <span>email: </span> <br />
                            <Input name="email" type="email" /> <br />
                        </div>
                    </div>

                    <div className="contentFornInsider">
                        <PhoneForm ></PhoneForm>
                    </div>
                    <div className="contentFornInsider">
                        <EnderecoForm></EnderecoForm>
                    </div>
                </div>
                <button type="submit">Cadastrar</button>
            </Form>
        );
    }
}

export default CadastroPessoa;