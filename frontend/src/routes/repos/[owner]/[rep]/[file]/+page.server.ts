import { error } from '@sveltejs/kit';
import fetch from 'node-fetch';


export function load({ params }) {
	const owner = params.owner;
	const rep = params.rep;
	const file = params.file;

	if (!rep) throw error(404);

	return {
		owner,
		rep,
		file,
	};
}