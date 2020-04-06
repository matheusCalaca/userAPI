import { Form } from '@unform/web';
import React, { Component } from 'react';
import EnderecoForm from '../../components/form/endereco-form';
import Input from '../../components/form/Input';
import PhoneForm from '../../components/form/phone-form';




class CadastroPessoa extends Component {
    handleSubmit = (data) => {
        console.log(JSON.stringify(data));
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

                    <div lassName="contentFornInsider">
                        <div>
                            <PhoneForm ></PhoneForm>
                        </div>
                    </div>
                    <div lassName="contentFornInsider">
                        <div>
                            <EnderecoForm></EnderecoForm>
                        </div>
                    </div>
                </div>
                <button type="submit">Cadastrar</button>
            </Form>
        );
    }
}

export default CadastroPessoa;