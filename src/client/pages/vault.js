import React from "react";
import Head from "next/head";
import Nav from "@/components/Nav";
import BottomNavbar from "@/components/BottomNavbar/BottomNavbar";
import Files from "@/components/Files/Files";

const Vault = () => {
    return (
        <>
            <Head>
                <title>Vault</title>
            </Head>
            <Nav />
            <Files />
            <BottomNavbar />
        </>
    );
};

export default Vault;
