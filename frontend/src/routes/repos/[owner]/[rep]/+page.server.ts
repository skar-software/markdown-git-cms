import { error } from '@sveltejs/kit';
import fetch from 'node-fetch';


export function load({ params }) {
	const rep = params.rep;
	const owner = params.owner;

	if (!rep) throw error(404);

	return {
		rep, owner
	};
}
