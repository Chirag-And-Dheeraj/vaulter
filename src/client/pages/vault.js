import React from "react";
import Head from "next/head";
import Nav from "@/components/Nav";
import SiteNav from "@/components/SiteNav/SiteNav";
import Files from "@/components/Files/Files";

const Vault = () => {
    return (
        <main className="md:grid md:grid-cols-12 md:grid-rows-4">
            <Head>
                <title>Vault</title>
            </Head>
            <Nav />
            <Files />
            <SiteNav />
        </main>
    );
};

export default Vault;
