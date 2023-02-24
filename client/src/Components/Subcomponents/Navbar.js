import React from 'react';
import './Navbar.css';
import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';

function Navbar1(props) {

    return (
            <Navbar bg="dark" variant="dark">
                <Container>
                    <Navbar.Brand href="/">Navbar</Navbar.Brand>
                    <Nav className="me-auto">
                        <Nav.Link href="/">Home</Nav.Link>
                        <Nav.Link href="/games">Games</Nav.Link>
                        <Nav.Link href="/characters">Characters</Nav.Link>
                    </Nav>
                </Container>
            </Navbar>
    );
}

export default Navbar1;