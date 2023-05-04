interface ResponseBodyType<T> {
    status: number,
    msg: string,
    data: T,
}

export default ResponseBodyType;