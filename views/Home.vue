<template>
	<div id="Home" class="columns">
		<div class="column is-4 is-offset-4">
			<b-field>
				<b-input placeholder="Search Videos"
					type="search"
					v-model="keyword"
					expanded>
				</b-input>
				<p class="control">
					<button class="button is-primary"
						@click="search">
						<font-awesome-icon icon="search" />
					</button>
				</p>
			</b-field>
		</div>
	</div>
</template>


<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { State, Action, Getter } from 'vuex-class';

const namespace: string = 'search';

@Component
export default class Home extends Vue {
	@Action('purgeVideos', { namespace }) purgeVideos: any;
	@Action('searchVideos', { namespace }) searchVideos: Promise<any>;

	public keyword: string = "";

	search() {
		if (this.keyword == "") {
			return
		}

		this.purgeVideos();

		this.searchVideos(this.keyword);
	}
}
</script>
