import Item from "./Item";

const Files = () => {
    return (
        <main className="p-3 md:col-start-4 md:col-end-13 lg:col-start-3">
            <Item name="Folder 1" lastModified="12 May 2022" size="122 MB" icon="files-filled" />
            <Item name="Folder 2" lastModified="1 May 2022" size="12 MB" icon="files-filled" />
            <Item name="File 1" lastModified="10 June 2022" size="12 KB" icon="document" />
            {/* <Item name="File 1" lastModified="10 June 2022" size="12 KB" icon="document" />
            <Item name="File 1" lastModified="10 June 2022" size="12 KB" icon="document" />
            <Item name="File 1" lastModified="10 June 2022" size="12 KB" icon="document" />
            <Item name="File 1" lastModified="10 June 2022" size="12 KB" icon="document" /> */}
        </main>
    );
};

export default Files;
