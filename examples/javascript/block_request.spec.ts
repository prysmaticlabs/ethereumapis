import {ListBlocksRequest} from 'examples_javascript/external/ethereumapis/eth/v1alpha1/beacon_chain_pb';


describe('ethereumapis', () => {
  it('allows creation of an object described by proto', () => {
    const req = new ListBlocksRequest();
    req.setSlot('555');

    expect(req.getSlot()).toEqual('555');
  });
});
