import React, {useEffect, useState} from 'react';
import {useParams} from "react-router-dom";
import useAlgoService from "../../services/UseAlgoService";
import {IAlgorithm} from "../../types/types";
import {marked} from "marked"
import  './TheoryBlock.scss'
import sanitizeHtml from 'sanitize-html';

const TheoryBlock = () => {
    const {algorithmId} = useParams();

    const {getAlgorithmTheory} = useAlgoService();
    const [algorithm, setAlgorithm] = useState<IAlgorithm>();



    //тестовый маркдаун
    const test_markdown =
        "### Очередь\n" +
        ">Структура данных, основанная на принципах FIFO - первый пришел, первый вышел. Есть две главные операции **push** (добавить в хвост) и **pop** (удалить элемент )\n" +
        "\n" +
        "![Очередь|100](https://neerc.ifmo.ru/wiki/images/thumb/c/c1/Fifo_new.png/225px-Fifo_new.png)\n" +
        "\n" +
        "Очередь поддерживает следующие операции:\n" +
        "-   **empty** - проверка очереди на наличие в ней элементов,\n" +
        "-   **push** - (запись в очередь) — операция вставки нового элемента,\n" +
        "-   **pop** - (снятие с очереди) — операция удаления нового элемента,\n" +
        "-   **size** - операция получения количества элементов в очереди.\n" +
        "\n" +
        "Идеи реализации очереди на массиве:\n" +
        "  1.  Добавлять в конец, перемещать все элемент на один вправо (долго)\n" +
        "  2. Хранить в двух переменных индексы начала и конца очереди\n" +
        "  3. Кольцевой буфер. Проблема: Надо знать макс ограничение элементов в очереди\n" +
        "\n" +
        "### Дек \n" +
        "deque - diuble ended queue\n" +
        "Эта структура поддерживает как и FIFO, так и LIFO.\n" +
        "На ее основе можно реализовать и очередь и стек.\n" +
        "\n" +
        "![deque](https://neerc.ifmo.ru/wiki/images/thumb/7/73/Deque1.png/300px-Deque1.png)\n" +
        "\n" +
        "Можно добавлять как в начало, так и в конец\n" +
        "У дека есть фронт(начало) и бек\n" +
        "Дека имеет следующие  операции:\n" +
        "-   empty — проверка на наличие элементов,\n" +
        "-   pushBack (запись в конец) — операция вставки нового элемента в конец,\n" +
        "-   popBack (снятие с конца) — операция удаления конечного элемента,\n" +
        "-   pushFront (запись в начало) — операция вставки нового элемента в начало,\n" +
        "-   popFront (снятие с начала) — операция удаления начального элемента.\n" +
        "\n" +
        "Можно реализовать на массиве, но лучше на двухсвязном списке. В случае двухсвязного списка, элемент хранит значение, предыдущего и следующий элемент\n" +
        "\n" +
        "Можно реализовать с помощью блоков\n" +
        "\n";


    useEffect(() => {
        getResources(Number(algorithmId) );
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [algorithmId]);

    const getResources = (algorithmId: number) => {
        getAlgorithmTheory(algorithmId)
            .then(onAlgorithmLoaded)
    }


    const onAlgorithmLoaded = (algorithm: IAlgorithm) => {
        setAlgorithm(algorithm);
        console.log(algorithm.content)
    }

    let dirty = marked(test_markdown);

    console.log(dirty)
    const clean = sanitizeHtml(dirty, {
        allowedTags: sanitizeHtml.defaults.allowedTags.concat([ 'img' ], ['blockquote']),
        allowedAttributes: false,
        selfClosing: [ 'img', 'br', 'hr', 'area', 'base', 'basefont', 'input', 'link', 'meta' ],
        allowedSchemes: [ 'http', 'https', 'ftp', 'mailto', 'tel' ],
        allowedSchemesByTag: {},
        allowedSchemesAppliedToAttributes: [ 'href', 'src', 'cite' ],
        allowProtocolRelative: true,
        enforceHtmlBoundary: false,
        parseStyleAttributes: true
    });

    return (
        <div className={'theory-content'} dangerouslySetInnerHTML ={{__html: clean}} />
    );
};

export default TheoryBlock;