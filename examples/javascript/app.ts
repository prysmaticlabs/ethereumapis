import {ChainHead} from 'examples_javascript/external/ethereumapis/eth/v1alpha1/beacon_chain_pb';

const head = new ChainHead();
head.setHeadSlot('555');

const el: HTMLDivElement = document.createElement('div');
el.innerText = `Chain head slot from server: ${head.getHeadSlot()}`;
el.className = 'ts1';
document.body.appendChild(el);
